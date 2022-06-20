package file

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFfmpegCommand(t *testing.T) {
	// -i: Input file (can be either an image or a video. If video, specify the starting frame)
	// -ss: Seek to the specified time (in seconds) in the input file
	// -vframes: Number of frames to extract
	// -vf: Video filter to use (scale)
	cmd := generateFfmpegCommand("/tmp/test.jpg", "/tmp/test.png", false, thumbnailWidth)
	assert.Equal(t,
		fmt.Sprintf("/usr/bin/ffmpeg -i /tmp/test.jpg -vframes 1 -vf scale=%d:-1 /tmp/test.png", thumbnailWidth),
		cmd.String())

	cmd = generateFfmpegCommand("/tmp/test.avi", "/tmp/test.png", true, thumbnailWidth)
	assert.Equal(t,
		fmt.Sprintf("/usr/bin/ffmpeg -i /tmp/test.avi -vframes 1 -vf scale=%d:-1 -ss 00:00:00 /tmp/test.png", thumbnailWidth),
		cmd.String())
}
