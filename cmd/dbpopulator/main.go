package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/CrowhopTech/shinysorter/backend/pkg/mongoimg"
	"github.com/CrowhopTech/shinysorter/backend/pkg/tickexecutor"
	"github.com/sirupsen/logrus"
)

var (
	imageMetadataConnection imagedb.ImageMetadata
)

var (
	importDirFlag       = flag.String("import-dir", "./import", "The directory to import files from")
	storageDirFlag      = flag.String("storage-dir", "./storage", "The directory to store files in")
	rescanIntervalFlag  = flag.Duration("rescan-interval", time.Second*5, "How often to rescan the import dir for new files")
	mongodbConectionURI = flag.String("mongodb-connection-uri", "mongodb://localhost:27017", "The connection URI for the MongoDB metadata database")
	purgeFlag           = flag.Bool("purge", false, "WARNING: If set to true, will empty the storage directory and reset the database")
)

func parseFlags() {
	flag.Parse()

	if result, err := os.Stat(*importDirFlag); err != nil {
		if os.IsNotExist(err) {
			logrus.Fatalf("Import directory '%s' does not exist: please create it and try again", *importDirFlag)
		} else {
			logrus.Fatalf("Error while checking info for import directory '%s'", *importDirFlag)
		}
	} else if !result.IsDir() {
		logrus.Fatalf("Import path '%s' exists but is not a directory", *importDirFlag)
	}

	if result, err := os.Stat(*storageDirFlag); err != nil {
		if os.IsNotExist(err) {
			logrus.Fatalf("Storage directory '%s' does not exist: please create it and try again", *storageDirFlag)
		} else {
			logrus.Fatalf("Error while checking info for storage directory '%s'", *storageDirFlag)
		}
	} else if !result.IsDir() {
		logrus.Fatalf("Storage path '%s' exists but is not a directory", *storageDirFlag)
	}

	return nil
}

func main() {
	rootCtx := context.Background()

	logrus.SetLevel(logrus.DebugLevel)

	parseFlags()

	// Initialize database connection
	mongoConn, cleanupFunc, err := mongoimg.New(rootCtx, *mongodbConectionURI, *purgeFlag)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize database connection")
	}
	defer cleanupFunc()

	imageMetadataConnection = mongoConn

	// Set ourselves up to receive system interrupts
	interruptSignals := make(chan os.Signal, 1)
	signal.Notify(interruptSignals, syscall.SIGINT, syscall.SIGTERM)

	cancelCtx, cancel := context.WithCancel(rootCtx)
	tickexecutor.New(cancelCtx, *rescanIntervalFlag, func(ctx context.Context) error {
		return scanForNewFiles(ctx, *importDirFlag, *storageDirFlag)
	}, nil)

	sig := <-interruptSignals
	cancel()
	logrus.Infof("Received interrupt '%s'", sig)
}

func scanForNewFiles(ctx context.Context, importDir string, storageDir string) error {
	logrus.Info("Doing scan!")

	wg := &sync.WaitGroup{}

	err := filepath.WalkDir(importDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		if d.IsDir() {
			// We don't process directories, just files
			return nil
		}

		go func() {
			processErr := processImportFile(ctx, wg, path, d, storageDir)
			if processErr != nil {
				logrus.WithError(processErr).WithField("file", path).Error("Failed to process file, leaving it behind to retry")
			}
		}()

		return nil
	})

	if err != nil {
		return err
	}

	wg.Wait()
	logrus.Infof("Done processing all files")

	return nil
}

func processImportFile(ctx context.Context, wg *sync.WaitGroup, path string, d fs.DirEntry, storageDir string) error {
	wg.Add(1)
	defer wg.Done()

	//   Check for conflicts by name. If exists, check hash. If match, continue. If not, warn and skip

	//   Copy to storage dir
	//   Add database entry
	//   Delete original entry

	finalTargetPath := filepath.Join(storageDir, d.Name())

	logrus.Infof("Processing file %s", path)

	inputFileMd5Sum, err := getFileMd5Sum(path)
	if err != nil {
		return err
	}

	exists, err := doesImageExist(finalTargetPath)
	if err != nil {
		return err
	}
	if exists {
		logrus.WithField("target_path", finalTargetPath).Debug("Target file already exists, checking file hash")
		existingFileMd5Sum, err := getFileMd5Sum(finalTargetPath)
		if err != nil {
			return err
		}
		if inputFileMd5Sum != existingFileMd5Sum {
			return fmt.Errorf("file already exists with name '%s' and hashes differ, skipping", d.Name())
		}
		logrus.WithField("target_path", finalTargetPath).Debug("Hashes match, skipping copy")
		// Continue since identical, just skip the copy. Try to create the entry: it'll fail if it already exists,
		// and will create if not.
	}

	if !exists {
		err = copyImageToStorageDir(path, finalTargetPath)
		if err != nil {
			return fmt.Errorf("failed to copy image '%s' to storage dir at '%s': %v", path, finalTargetPath, err)
		}
	}

	err = createOrCheckEntryForFile(ctx, storageDir, d.Name())
	if err != nil {
		return fmt.Errorf("failed to create entry for file '%s': %v", d.Name(), err)
	}

	err = deleteFile(path)
	if err != nil {
		logrus.WithError(err).
			WithField("input_path", path).
			Warn("Failed to delete input file, it may start throwing duplicate errors")
	}

	return nil
}

// Filenames
func doesImageExist(destFilePath string) (bool, error) {
	_, err := os.Stat(destFilePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// Takes in an absolute path to src file and absolute path to dest file
func copyImageToStorageDir(srcFilePath string, destFilePath string) error {
	srcFileReader, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFileReader.Close()

	destFileWriter, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFileWriter.Close()

	bytesCopied, err := io.Copy(destFileWriter, srcFileReader)
	if err != nil {
		return err
	}
	if bytesCopied == 0 {
		return fmt.Errorf("zero bytes copied")
	}
	return nil
}

// Filename
func createOrCheckEntryForFile(ctx context.Context, storageDir string, fileName string) error {
	img, err := newImageEntry(ctx, storageDir, fileName)
	if err != nil {
		return err
	}

	return imageMetadataConnection.CreateImageEntry(ctx, &img)
}

func newImageEntry(ctx context.Context, storageDir string, fileName string) (imagedb.Image, error) {
	md5Sum, err := getFileMd5Sum(filepath.Join(storageDir, fileName))
	if err != nil {
		return imagedb.Image{}, err
	}

	imageEntry := imagedb.Image{
		FileMetadata: imagedb.FileMetadata{
			Name:   fileName,
			Md5Sum: md5Sum,
		},
	}
	return imageEntry, nil
}

// OS path
func deleteFile(path string) error {
	err := os.Remove(path)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// OS path
func getFileMd5Sum(path string) (string, error) {
	md5Summer := md5.New()
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	writtenBytes, err := io.Copy(md5Summer, file)
	if err != nil {
		return "", err
	}

	if writtenBytes == 0 {
		return "", fmt.Errorf("zero bytes copied while summing file")
	}

	return fmt.Sprintf("%x", md5Summer.Sum(nil)), nil
}
