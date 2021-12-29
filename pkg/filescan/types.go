package filescan

import "context"

type FileListener func(ctx context.Context, f *FileEntry) error

type FileScanner interface {
	// RegisterCallback will set the given listener to be
	// called whenever a Refresh is triggered on a file or
	// the Watch finds a new/updated entry.
	RegisterCallback(fl FileListener)

	// Watch will block the goroutine and will check for new files and updated files (metadata or contents).
	// Removed files are not yet supported.
	// If stopOnScanError is true, errors in listing files will cause the Watch to stop and return an error.
	// if stopOnCallbackError is true, errors in the triggered callback will cause the Watch to stop and return the error.
	Watch(ctx context.Context, stopOnScanError bool, stopOnCallbackError bool) error

	// Refresh will trigger sending the given file to the callback.
	// Useful if a file got out of sync or you're doing database ops.
	Refresh(ctx context.Context, f *FileEntry) error
}
