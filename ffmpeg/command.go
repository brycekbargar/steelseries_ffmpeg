package ffmpeg

import (
	"fmt"
	"io"
	"time"
)

// FilePath is a path to a file on disk.
type FilePath string

// Command contains context about a video to run through ffmpeg.
type Command struct {
	// Input is the location of the video file on disk.
	Input FilePath
	// Output is the location where ffmpeg should output the video resulting from the command.
	Output FilePath
	// Width determines the video's size.
	Width uint
	// Height determines the video's size.
	Height uint
	// Duration is the length of the video.
	Duration time.Duration
}

// NewCommand creates (and todo validates) the basic metadata for an ffmpeg command.
func NewCommand(
	input FilePath,
	output FilePath,
	resolution struct{ X, Y uint },
	duration string) (*Command, error) {

	d, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}

	c := &Command{
		Input:    input,
		Output:   output,
		Width:    resolution.X,
		Height:   resolution.Y,
		Duration: d,
	}

	// return Validate(c)
	return c, nil
}

// Render writes the ffmpeg command string to the io.Writer.
func (c Command) Render(out io.Writer) error {
	_, err := fmt.Fprintf(out, "ffmpeg -i %v ", c.Input)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%v", c.Output)
	if err != nil {
		return err
	}

	return nil
}
