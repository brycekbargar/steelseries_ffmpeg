package ffmpeg

import (
	"fmt"
	"io"
	"time"
)

// DrawtextFilter represents the collection of parameters necessary to add a an ffmpeg 'drawtext' filter to a video.
type DrawtextFilter struct {
	Value             string
	X, Y, Size, Color uint
	Start, End        float64
}

// AddDrawtextFilter creates a valid ffmpeg 'drawtext' filter and adds it to the command.
func (c *Command) AddDrawtextFilter(
	value string,
	location struct{ X, Y uint },
	size uint,
	color uint,
	start string,
	end string,
) error {

	s, err := time.ParseDuration(start)
	if err != nil {
		return err
	}
	e, err := time.ParseDuration(end)
	if err != nil {
		return err
	}

	f := &DrawtextFilter{
		Value: value,
		X:     location.X,
		Y:     location.Y,
		Size:  size,
		Color: color,
		Start: s.Seconds(),
		End:   e.Seconds(),
	}
	if f.Start > c.Duration.Seconds() {
		return ErrInvalidFilterStartTime
	}
	if f.End > c.Duration.Seconds() {
		return ErrInvalidFilterEndTime
	}
	if f.X > c.Width || f.Y > c.Height {
		return ErrInvalidFilterLocation
	}

	c.Filters = append(c.Filters, f)
	return nil
}

// Render writes the ffmpeg 'drawtext' filter string to the io.Writer.
func (f DrawtextFilter) Render(out io.Writer) error {
	_, err := fmt.Fprintf(
		out,
		"drawtext = enable='between(t,%.1f,%.1f)':text='%v':fontcolor=0x%06X:fontsize=%d:x=%d:y=%d",
		f.Start,
		f.End,
		f.Value,
		f.Color,
		f.Size,
		f.X,
		f.Y,
	)
	if err != nil {
		return err
	}
	return nil
}
