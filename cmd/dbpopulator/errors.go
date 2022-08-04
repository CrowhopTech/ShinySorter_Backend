package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const errorsFileName = "import_errors.txt"

var (
	importErrors = map[string]error{}
	errorsLock   = &sync.Mutex{}
)

func writeErrors(baseDir string) error {
	errorsLock.Lock()
	defer errorsLock.Unlock()

	sb := strings.Builder{}
	for file, err := range importErrors {
		absolutePath := path.Join(baseDir, file)
		logrus.Debugf("Checking if path '%s' still exists for errors", absolutePath)
		_, statErr := os.Stat(absolutePath)
		if statErr != nil {
			if !os.IsNotExist(statErr) {
				logrus.Debugf("File '%s' did not exist, clearing errors", absolutePath)
				return statErr
			}
			delete(importErrors, file)
			continue
		}
		sb.WriteString(fmt.Sprintf("'%s': %v\n", file, err))
	}
	os.WriteFile(path.Join(baseDir, errorsFileName), []byte(sb.String()), os.FileMode(0644))

	return nil
}

func setErrorForFile(filename string, fileErr error) bool {
	errorsLock.Lock()
	defer errorsLock.Unlock()

	previousValue := ""
	if err, ok := importErrors[filename]; ok && err != nil {
		previousValue = err.Error()
	}

	importErrors[filename] = fileErr
	return fileErr.Error() != previousValue
}

func clearErrorForFile(filename string) bool {
	errorsLock.Lock()
	defer errorsLock.Unlock()

	previouslyHadValue := false
	if err, ok := importErrors[filename]; ok && err != nil {
		previouslyHadValue = true
	}

	delete(importErrors, filename)
	return previouslyHadValue // Basically returns if we need to write the change or not
}

func setErrorForFileAndWrite(baseDir string, filename string, fileErr error) error {
	requiresChange := setErrorForFile(filename, fileErr)
	if requiresChange {
		return writeErrors(baseDir)
	}
	return nil
}

func clearErrorForFileAndWrite(baseDir string, filename string) error {
	requiresChange := clearErrorForFile(filename)
	if requiresChange {
		return writeErrors(baseDir)
	}
	return nil
}
