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
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/CrowhopTech/shinysorter/backend/pkg/tickexecutor"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	apiclient "github.com/CrowhopTech/shinysorter/backend/pkg/swagger/client"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/client/operations"
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
		return scanForNewFiles(ctx, *importDirFlag)
	}, nil)

	sig := <-interruptSignals
	cancel()
	logrus.Infof("Received interrupt '%s'", sig)
}

func scanForNewFiles(ctx context.Context, importDir string) error {
	logrus.Debug("Testing if API server is reachable...")
	_, err := swaggerClient.Operations.CheckHealth(operations.NewCheckHealthParams())
	if err != nil {
		return fmt.Errorf("API server is not accessible, skipping scan: %v", err)
	}
	logrus.Debug("API server accessible, doing scan!")

	wg := &sync.WaitGroup{}

	err = filepath.WalkDir(importDir, func(path string, d fs.DirEntry, err error) error {
		logFields := logrus.WithField("file", path)

		if err != nil {
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

		trimmedPath := strings.TrimPrefix(path, importDir)

		if checkIfFileLocked(trimmedPath) {
			logFields.WithField("trimmedPath", trimmedPath).Debug("A scan routine is already running for file, not starting a new one")
			return nil
		}

		go func() {
			lockFile(trimmedPath)
			defer unlockFile(trimmedPath)
			processErr := processImportFile(ctx, wg, path, d)
			if processErr != nil {
				logFields.WithError(processErr).Error("Failed to process file, leaving it behind to retry")
				writeErr := setErrorForFileAndWrite(importDir, trimmedPath, processErr)
				if writeErr != nil {
					logFields.WithError(writeErr).Error("Failed to write errors to file")
				}
			} else {
				writeErr := clearErrorForFileAndWrite(importDir, trimmedPath)
				if writeErr != nil {
					logFields.WithError(writeErr).Error("Failed to write errors to file")
				}
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

func waitForFileToStabilize(ctx context.Context, path string, d fs.DirEntry) error {
	lastSize := int64(0)
	for {
		time.Sleep(fileSizeDebounceTime) // Wait a bit to give the file time to finish writing

		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}

		// Read the file size and compare to the last loop
		fileInfo, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("failed to stat file '%s': %v", d.Name(), err)
		}

		currentSize := fileInfo.Size()
		if currentSize != lastSize {
			if lastSize == int64(0) {
				logrus.Debugf("First file scan for file '%s', size is %d", d.Name(), currentSize)
			} else {
				logrus.Infof("File '%s' is still being written to, now %d bytes", d.Name(), currentSize)
			}
			lastSize = currentSize
			continue
		}

		// File hasn't changed size since the last check: it's probably done writing
		return nil
	}
}

func processImportFile(ctx context.Context, wg *sync.WaitGroup, path string, d fs.DirEntry) error {
	wg.Add(1)
	defer func() {
		wg.Done()
	}()
	// Wait until file is done being written to
	// Create file entry with file hash set (will fail if doesn't match)
	// Set file contents
	// Delete original file

	logrus.Infof("Waiting for file %s to stabilize", path)
	err := waitForFileToStabilize(ctx, path, d)
	if err != nil {
		logrus.Errorf("Failed: %v", err)
		return err
	}

	logrus.Infof("Processing file %s", path)

	err = createOrCheckEntryForFile(ctx, path)
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

// Filename
func createOrCheckEntryForFile(ctx context.Context, importFile string) error {
	img, err := newFileEntry(ctx, importFile)
	if err != nil {
		return err
	}

	// Will also fail if the md5sum deosn't match (how we check for conflicts)
	_, err = swaggerClient.Operations.CreateFile(operations.NewCreateFileParams().WithNewFile(&img))
	if err != nil {
		return fmt.Errorf("failed to create file through REST: %v", err)
	}

	file, err := os.Open(importFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	_, err = swaggerClient.Operations.SetFileContent(operations.NewSetFileContentParams().
		WithContext(ctx).
		WithID(img.ID).
		WithFileContents(file),
	)
	if err != nil {
		return fmt.Errorf("failed to set file contents through REST: %v", err)
	}
	logrus.Infof("finished sending")

	return nil
}

func newFileEntry(ctx context.Context, filePath string) (models.File, error) {
	md5Sum, err := getFileMd5Sum(filePath)
	if err != nil {
		return models.File{}, err
	}
	f := false
	return models.File{
		ID:            filepath.Base(filePath),
		Md5sum:        md5Sum,
		HasBeenTagged: &f,
	}, nil
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
