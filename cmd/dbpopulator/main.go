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
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/CrowhopTech/shinysorter/backend/pkg/tickexecutor"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	apiclient "github.com/CrowhopTech/shinysorter/backend/pkg/swagger/client"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/client/files"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/models"
	httptransport "github.com/go-openapi/runtime/client"
)

const (
	fileSizeDebounceTime = time.Second * 5
)

var (
	importDirFlag      = flag.String("import-dir", "./import", "The directory to import files from")
	rescanIntervalFlag = flag.Duration("rescan-interval", time.Second*5, "How often to rescan the import dir for new files")
	restAddressFlag    = flag.String("rest-address", "localhost:10000", "The address (and port, no protocol) to reach the REST server")
	logLevel           = flag.String("log-level", "info", "The log level to print at")

	swaggerClient *apiclient.ShinySorter
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

	parsedLevel, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		logrus.Panicf("Failed to parse log level %s", *logLevel)
	}
	logrus.SetLevel(parsedLevel)
}

func main() {
	rootCtx := context.Background()

	parseFlags()

	logrus.Info("Constructing Swagger client...")

	swaggerClient = apiclient.New(
		httptransport.New(*restAddressFlag, "/", []string{"http"}),
		strfmt.Default,
	)

	// Set ourselves up to receive system interrupts
	interruptSignals := make(chan os.Signal, 1)
	signal.Notify(interruptSignals, syscall.SIGINT, syscall.SIGTERM)

	cancelCtx, cancel := context.WithCancel(rootCtx)
	logrus.Infof("Starting interval scanner with interval %v", *rescanIntervalFlag)
	tickexecutor.New(cancelCtx, *rescanIntervalFlag, func(ctx context.Context) error {
		return scanForNewFiles(ctx, path.Clean(*importDirFlag))
	}, nil)

	go backgroundProcessFiles(cancelCtx, *importDirFlag)

	sig := <-interruptSignals
	cancel()
	logrus.Infof("Received interrupt '%s'", sig)
}

func scanForNewFiles(ctx context.Context, importDir string) error {
	err := filepath.WalkDir(importDir, func(filePath string, d fs.DirEntry, err error) error {
		// Give a chance to break for context
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}
		if err != nil {
			// Some error in actually reading the files
			return filepath.SkipDir
		}
		if d.IsDir() {
			// We don't process directories, just files
			return nil
		}
		if d.Name() == errorsFileName {
			// Don't process the file holding the error messages, since it's auto-generated
			return nil
		}

		trimmedPath := strings.TrimPrefix(filePath, importDir+"/")
		logrus.WithField("trimmed_file", trimmedPath).Debug("Ensuring file is tracked")
		ensureFileTracked(trimmedPath)

		return nil
	})

	if err != nil {
		return err
	}
	logrus.Infof("Done processing all files")

	return nil
}

// backgroundProcessFiles will run continually until the context is cancelled, picking up one file at a time
// and processing it
func backgroundProcessFiles(ctx context.Context, importDir string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		nextFile := getNextFile()
		if nextFile == "" {
			time.Sleep(time.Second)
			continue
		}
		err := processFile(ctx, importDir, nextFile)
		if err != nil {
			if err != nil {
				writeErr := setErrorForFileAndWrite(importDir, nextFile, err)
				if writeErr != nil {
					logrus.WithField("file", nextFile).WithError(writeErr).Error("Failed to write errors to file")
				}
			} else {
				writeErr := clearErrorForFileAndWrite(importDir, nextFile)
				if writeErr != nil {
					logrus.WithField("file", nextFile).WithError(writeErr).Error("Failed to write errors to file (while clearing an error)")
				}
			}
		}
	}
}

// processFile is the entire import routine for one file, and will release the file at the end
func processFile(ctx context.Context, importDir string, file string) error {
	defer completeFile(file)

	logFields := logrus.WithField("file", file)

	err := debounceFileSize(ctx, importDir, file)
	if err != nil {
		logFields.WithError(err).Error("Error while waiting for file size to settle")
		return fmt.Errorf("error while waiting for file size to settle: %v", err)
	}

	err = createFileEntryOnServer(ctx, importDir, file)
	if err != nil {
		logFields.WithError(err).Error("Error while creating entry for file on server")
		return fmt.Errorf("error while creating entry for file on server: %v", err)
	}

	err = uploadFileContents(ctx, importDir, file)
	if err != nil {
		logFields.WithError(err).Error("Error while uploading file contents")
		return fmt.Errorf("error while uploading file contents: %v", err)
	}

	err = cleanUpFile(ctx, importDir, file)
	if err != nil {
		logFields.WithError(err).Error("Error while cleaning up file")
		return fmt.Errorf("error while cleaning up file: %v", err)
	}

	logFields.Info("Successfully processed file")
	return nil
}

func getFileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func debounceFileSize(ctx context.Context, importDir string, file string) error {
	fullPath := path.Join(importDir, file)
	lastSize, err := getFileSize(fullPath)
	if err != nil {
		return fmt.Errorf("failed to stat file '%s': %v", fullPath, err)
	}

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}

		time.Sleep(fileSizeDebounceTime) // Wait a bit to give the file time to finish writing

		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}

		// Read the file size and compare to the last loop
		currentSize, err := getFileSize(fullPath)
		if err != nil {
			return fmt.Errorf("failed to stat file '%s': %v", fullPath, err)
		}

		if currentSize == lastSize {
			// File hasn't changed size since the last check: it's probably done writing
			return nil
		}
		logrus.Infof("File '%s' is still being written to, now %d bytes", file, currentSize)
		lastSize = currentSize
	}
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

func createFileEntryOnServer(ctx context.Context, importDir string, file string) error {
	md5Sum, err := getFileMd5Sum(path.Join(importDir, file))
	if err != nil {
		return fmt.Errorf("failed to get file md5sum: %v", err)
	}
	f := false
	img := models.File{
		ID:            filepath.Base(file),
		Md5sum:        md5Sum,
		HasBeenTagged: &f,
	}

	// Create entry on the server
	// Will also fail if the md5sum deosn't match (this is how we check for conflicts)
	_, err = swaggerClient.Files.CreateFile(files.NewCreateFileParams().WithNewFile(&img))
	if err != nil {
		return fmt.Errorf("failed to create file through REST: %v", err)
	}

	return nil
}

func uploadFileContents(ctx context.Context, importDir string, file string) error {
	osFile, err := os.Open(path.Join(importDir, file))
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	_, err = swaggerClient.Files.SetFileContent(files.NewSetFileContentParams().
		WithContext(ctx).
		WithID(filepath.Base(file)).
		WithFileContents(osFile),
	)
	if err != nil {
		return fmt.Errorf("failed to set file contents through REST: %v", err)
	}
	return nil
}

func cleanUpFile(ctx context.Context, importDir string, file string) error {
	err := os.Remove(path.Join(importDir, file))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
