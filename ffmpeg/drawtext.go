package ffmpeg

import (
	"fmt"
	"io"
	"time"
)

type DrawTextFilter struct {
	Value             string
	X, Y, Size, Color uint
	Start, End        float64
}

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

	f := &DrawTextFilter{
		Value: value,
		X:     location.X,
		Y:     location.Y,
		Size:  size,
		Color: color,
		Start: s.Seconds(),
		End:   e.Seconds(),
	}
	// Validate(f)

	c.Filters = append(c.Filters, f)
	return nil
}

func (f DrawTextFilter) Render(out io.Writer) error {
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
