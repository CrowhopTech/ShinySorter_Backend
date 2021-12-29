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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/CrowhopTech/shinysorter/backend/pkg/tickexecutor"
	"github.com/sirupsen/logrus"
)

const (
	databaseName         = "shiny_sorter"
	imagesCollectionName = "images"
)

var (
	imagesCollection *mongo.Collection
)

type Image struct {
	Name   string `bson:"_id"`
	Md5Sum string `bson:"md5sum"`
}

func main() {
	rootCtx := context.Background()

	logrus.SetLevel(logrus.DebugLevel)

	importDirFlag := flag.String("import-dir", "./import", "The directory to import files from")
	storageDirFlag := flag.String("storage-dir", "./storage", "The directory to store files in")
	rescanIntervalFlag := flag.Duration("rescan-interval", time.Second*5, "How often to rescan the import dir for new files")
	mongodbConectionURI := flag.String("mongodb-connection-uri", "mongodb://localhost:27017", "The connection URI for the MongoDB metadata database")
	purgeFlag := flag.Bool("purge", false, "WARNING: If set to true, will empty the storage directory and reset the database")

	flag.Parse()

	if _, err := os.Stat(*importDirFlag); err != nil {
		if os.IsNotExist(err) {
			logrus.Fatalf("Import directory '%s' does not exist: please create it and try again", *importDirFlag)
		} else {
			logrus.Fatalf("Error while checking info for import directory '%s'", *importDirFlag)
		}
	}

	if _, err := os.Stat(*storageDirFlag); err != nil {
		if os.IsNotExist(err) {
			logrus.Fatalf("Storage directory '%s' does not exist: please create it and try again", *storageDirFlag)
		} else {
			logrus.Fatalf("Error while checking info for storage directory '%s'", *storageDirFlag)
		}
	}

	cleanupFunc, err := initDB(rootCtx, *purgeFlag, *mongodbConectionURI)
	if cleanupFunc != nil {
		defer func() {
			cleanupErr := cleanupFunc()
			if cleanupErr != nil {
				logrus.WithError(cleanupErr).Error("Failed to clean up database connection")
			}
		}()
	}

	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize database connection")
	}

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

func initDB(ctx context.Context, purge bool, uri string) (func() error, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	cleanupFunc := func() error {
		return client.Disconnect(context.TODO())
	}

	if purge {
		err = client.Database(databaseName).Drop(ctx)
		if err != nil {
			return nil, err
		}
	}

	imagesCollection = client.Database(databaseName).Collection(imagesCollectionName)

	return cleanupFunc, nil
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
	// TODO: calculate hash
	img, err := newImageEntry(ctx, storageDir, fileName)
	if err != nil {
		return err
	}

	// Find the existing entry. If we didn't, insert one.
	res := imagesCollection.FindOne(ctx, bson.M{
		"_id": img.Name,
	})
	if res.Err() == nil {
		// Entry already existed, validate that it matches the entry we're uploading
		existingEntry := Image{}
		err = res.Decode(&existingEntry)
		if err != nil {
			return err
		}

		// Check all relevant parameters: if they all match, DO NOTHING (don't overwrite tags or anything)
		if existingEntry.Md5Sum == img.Md5Sum {
			return nil
		}

		// Entry already exists but it doesn't match up with this file: do nothing and wait for user intervention
		return fmt.Errorf("database entry already exists for file '%s' and is different", img.Name)
	}

	if res.Err() != mongo.ErrNoDocuments {
		return res.Err()
	}

	_, err = imagesCollection.InsertOne(ctx, img)
	if err != nil {
		return err
	}
	return nil
}

func newImageEntry(ctx context.Context, storageDir string, fileName string) (Image, error) {
	md5Sum, err := getFileMd5Sum(filepath.Join(storageDir, fileName))
	if err != nil {
		return Image{}, err
	}

	imageEntry := Image{
		Name:   fileName,
		Md5Sum: md5Sum,
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
