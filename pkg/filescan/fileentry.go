package filescan

import (
	"time"

	"github.com/sirupsen/logrus"
)

type FileEntry struct {
	Path  string
	ADate time.Time
	Inode string
}

func (f *FileEntry) LogFields(verbose bool) *logrus.Entry {
	fields := logrus.WithFields(logrus.Fields{
		"inode":      f.Inode,
		"path":       f.Path,
		"changed_at": f.ADate,
	})
	if verbose {
		// As we get more information, the less important more debuggy stuff can go here
		fields = fields.WithFields(logrus.Fields{})
	}

	return fields
}
