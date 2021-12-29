package filescan

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	findParseFormat = "2006-01-02 3:04:05 PM MST"
	findPrintFormat = "%p\\t%AY-%Am-%Ad %Ar %AZ\\t%i\\n"
)

var _ FileScanner = new(adateScanner)

// adateScanner is an implementation of FileScanner that will
// repeatedly check for files accessed after the end of the last
// check. This is horrendously inefficient especially considering
// that we may access these files as part of the scan process.
type adateScanner struct {
	fileRescanInterval time.Duration
	searchPath         string
	rescanWall         time.Time
	updateLatestTime   func(time.Time) error
	callback           FileListener
}

func NewADateScanner(fileRescanInterval time.Duration, searchPath string, startTime *time.Time, updateLatestTime func(time.Time) error) (FileScanner, error) {
	// TODO: check that search path exists?
	var rescanWall time.Time
	if startTime != nil {
		rescanWall = *startTime
	}
	return &adateScanner{
		fileRescanInterval: fileRescanInterval,
		searchPath:         searchPath,
		rescanWall:         rescanWall,
		updateLatestTime:   updateLatestTime,
	}, nil
}

// RegisterCallback will set the given listener to be
// called whenever a Refresh is triggered on a file or
// the Watch finds a new/updated entry.
func (ads *adateScanner) RegisterCallback(fl FileListener) {
	if ads.callback != nil {
		panic("adateScanner callback has already been registered")
	}
	ads.callback = fl
}

// Watch should be called in a goroutine and will check
// for new files and updated files (metadata or contents).
// Removed files are not yet supported.
// If stopOnScanError is true, errors in listing files will cause the Watch to stop and return an error.
// if stopOnCallbackError is true, errors in the triggered callback will cause the Watch to stop and return the error.
func (ads *adateScanner) Watch(ctx context.Context, stopOnScanError bool, stopOnCallbackError bool) error {
	t := time.NewTicker(ads.fileRescanInterval)
	for {
		// Get files
		// Iterate them for callbacks
		// Update time
		// Wait on channel

		scanTime := ads.rescanWall.Add(time.Second)
		newFiles, err := ads.findNewFiles(&scanTime)
		if err != nil {
			if stopOnScanError {
				logrus.WithError(err).Error("Failed to scan for new files")
				return err
			} else {
				logrus.WithError(err).Warn("Failed to scan for new files, waiting and retrying...")
				<-t.C
				continue
			}
		}

		newestAdate := ads.rescanWall

		failOut := false

		for _, f := range newFiles {
			if f.ADate.After(newestAdate) {
				newestAdate = f.ADate
			}

			if ads.callback != nil {
				// Invoke callback, process errors if they come up
				err = ads.callback(ctx, f)
				if err != nil {
					if stopOnCallbackError {
						f.LogFields(false).WithError(err).Error("Failed to process new files")
						return err
					} else {
						f.LogFields(false).WithError(err).Warn("Failed to process new files, waiting and retrying...")
						failOut = true
						break
					}
				}
			}
		}

		if failOut {
			// Do the wait for next loop
			<-t.C
			continue
		}
		ads.rescanWall = newestAdate
		err = ads.updateLatestTime(ads.rescanWall)
		if err != nil {
			logrus.WithError(err).Warn("Failed to update latest time, may waste effort rescanning in the future")
		}

		logrus.WithField("sleep_time", ads.fileRescanInterval).Debug("Scan complete, waiting...")

		<-t.C
	}
}

// Best effort: we'll do more integrity-based scans later, I just want to unblock myself
// and fidgeting around with birth dates and ctimes is making my head spin
func (ads *adateScanner) findNewFiles(modifiedAfter *time.Time) ([]*FileEntry, error) {
	// TODO: "page" this with max count and sort by mtime so that we don't revisit.
	//       Store mtime of latest file, or if less than max count, last page (?)
	//		 Then break with a "retry fast? y/n" flag

	args := []string{
		ads.searchPath,
		"-mindepth", "1",
		"-type", "f",
	}

	if modifiedAfter != nil {
		args = append(args, "-newerat", modifiedAfter.Format(findParseFormat))
	}

	args = append(args, "-printf", findPrintFormat)

	cmd := exec.Command("find", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("find failed: %v (%s)", err, string(out))
	}
	lines := strings.Split(string(out), "\n")
	var newFiles []*FileEntry
	for _, line := range lines {
		parsed, err := parseFindLine(line)
		if err != nil {
			logrus.WithError(err).WithField("line", line).Warn("Find entry is invalid, skipping")
			continue
		}
		if parsed == nil {
			continue
		}

		parsed.LogFields(false).Debug("Found file")
		newFiles = append(newFiles, parsed)
	}

	return newFiles, nil
}

func parseFindLine(line string) (*FileEntry, error) {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil, nil
	}

	components := strings.Split(line, "\t")
	if len(components) != 3 {
		return nil, fmt.Errorf("line has wrong number of components %d, expected 3", len(components))
	}

	path := components[0]
	dateRaw := components[1]
	inodeRaw := components[2]

	date, err := time.Parse(findParseFormat, dateRaw)
	if err != nil {
		return nil, fmt.Errorf("date '%s' is invalid: %v", dateRaw, err)
	}

	inode, err := strconv.ParseUint(inodeRaw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("inode '%s' is invalid: %v", inodeRaw, err)
	}
	return &FileEntry{
		Path:  path,
		Inode: strconv.FormatUint(inode, 10),
		ADate: date,
	}, nil
}

// Refresh will trigger sending the given file to the callback.
// Useful if a file got out of sync or you're doing database ops.
func (ads *adateScanner) Refresh(ctx context.Context, f *FileEntry) error {
	if ads.callback == nil {
		return nil
	}
	return ads.callback(ctx, f)
}
