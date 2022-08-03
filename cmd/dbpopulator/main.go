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

	"github.com/CrowhopTech/shinysorter/backend/pkg/tickexecutor"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"

	apiclient "github.com/CrowhopTech/shinysorter/backend/pkg/swagger/client"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/client/operations"
	"github.com/CrowhopTech/shinysorter/backend/pkg/swagger/models"
	httptransport "github.com/go-openapi/runtime/client"
)

var (
	swaggerClient *apiclient.ShinySorter
)

var (
	importDirFlag      = flag.String("import-dir", "./import", "The directory to import files from")
	rescanIntervalFlag = flag.Duration("rescan-interval", time.Second*5, "How often to rescan the import dir for new files")
	restAddressFlag    = flag.String("rest-address", "localhost:10000", "The address (and port, no protocol) to reach the REST server")
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
}

func main() {
	rootCtx := context.Background()

	logrus.SetLevel(logrus.InfoLevel)

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
	logrus.Debug("Doing scan!")

	// TODO: issue #13 (locking around files)

	logrus.Debug("Testing if API server is reachable...")
	_, err := swaggerClient.Operations.CheckHealth(operations.NewCheckHealthParams())
	if err != nil {
		return fmt.Errorf("API server is not accessible, skipping scan: %v", err)
	}
	logrus.Debug("API server accessible")

	wg := &sync.WaitGroup{}

	err = filepath.WalkDir(importDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		if d.IsDir() {
			// We don't process directories, just files
			return nil
		}

		go func() {
			processErr := processImportFile(ctx, wg, path, d)
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

func processImportFile(ctx context.Context, wg *sync.WaitGroup, path string, d fs.DirEntry) error {
	wg.Add(1)
	defer wg.Done()

	// Create file entry with file hash set (will fail if doesn't match)
	// Set file contents
	// Delete original file

	logrus.Infof("Processing file %s", path)

	err := createOrCheckEntryForFile(ctx, path)
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
