package main

import "sync"

var (
	currentFilesLock sync.Mutex      = sync.Mutex{}
	currentFiles     map[string]bool = make(map[string]bool)
)

func checkIfFileLocked(file string) bool {
	currentFilesLock.Lock()
	defer currentFilesLock.Unlock()

	_, ok := currentFiles[file]
	return ok
}

func lockFile(file string) {
	currentFilesLock.Lock()
	defer currentFilesLock.Unlock()

	currentFiles[file] = true
}

func unlockFile(file string) {
	currentFilesLock.Lock()
	defer currentFilesLock.Unlock()

	delete(currentFiles, file)
}
