package file

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/h2non/filetype/types"
)

const thumbnailWidth = 640

func WriteFileThumbnail(path string, ft types.Type, outputPath string) error {
	if strings.HasPrefix(ft.MIME.Type, "image") {
		return runCommandWrapper(generateFfmpegCommand(path, outputPath, false, thumbnailWidth))
	}
	if strings.HasPrefix(ft.MIME.Type, "video") {
		return runCommandWrapper(generateFfmpegCommand(path, outputPath, true, thumbnailWidth))
	}
	return fmt.Errorf("unsupported MIME type to make thumbnail for path '%s': %s", path, ft.MIME.Type)
}

func generateFfmpegCommand(inputPath string, outputPath string, isVideo bool, targetWidth uint) *exec.Cmd {
	// -i: Input file (can be either an image or a video. If video, specify the starting frame)
	// -ss: Seek to the specified time (in seconds) in the input file
	// -vframes: Number of frames to extract
	// -vf: Video filter to use (scale)
	args := []string{"-i", inputPath, "-vframes", "1", "-vf", fmt.Sprintf("scale=%d:-1", targetWidth)}
	if isVideo {
		args = append(args, "-ss", "00:00:00")
	}
	args = append(args, outputPath)
	return exec.Command("ffmpeg", args...)
}

func runCommandWrapper(cmd *exec.Cmd) error {
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error while running command '%v': %v, stderr: %s", cmd, err, stderr.String())
	}

	return nil
}
