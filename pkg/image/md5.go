package image

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// GetFileMd5Sum takes in an OS file path and returns the md5sum as a string
func GetFileMd5Sum(path string) (string, error) {
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
