package main

import "sync"

var (
	fileQueueLock sync.Mutex      = sync.Mutex{}
	fileQueue     []string        = []string{}
	allFiles      map[string]bool = make(map[string]bool)
)

// Called to add a file to the processing queue if it's not already being handled
func ensureFileTracked(file string) {
	fileQueueLock.Lock()
	defer fileQueueLock.Unlock()

	if _, ok := allFiles[file]; !ok {
		fileQueue = append(fileQueue, file)
		allFiles[file] = true
	}
}

// Called when the thread pulls something from the queue
func getNextFile() string {
	fileQueueLock.Lock()
	defer fileQueueLock.Unlock()

	if len(fileQueue) == 0 {
		return ""
	}

	nextFile := fileQueue[0]
	fileQueue = fileQueue[1:] // Remove first element

	allFiles[nextFile] = true // Just to be safe, make sure it's still tracked in the map

	return nextFile
}

// Called once a thread is done with a file, whether success or error
func completeFile(file string) {
	fileQueueLock.Lock()
	defer fileQueueLock.Unlock()

	delete(allFiles, file)
}
