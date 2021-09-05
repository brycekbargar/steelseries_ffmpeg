package ffmpeg_test

import (
	"testing"

	"github.com/brycekbargar/steelseries_ffmpeg/ffmpeg"
	"github.com/stretchr/testify/assert"
)

func TestCommand_AddDrawtextFilter_Invalid(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		command  *ffmpeg.Command
		value    string
		location struct{ X, Y uint }
		size     uint
		color    uint
		start    string
		end      string
		err      error
	}{
		{
			"Filter Test #3",
			&ffmpeg.Command{
				Duration: MustParseDuration("60.0s"),
			},
			"RIP",
			struct{ X, Y uint }{100, 200},
			32,
			0xFFFFFF,
			"24.0s",
			"75.0s",
			ffmpeg.ErrInvalidFilterEndTime,
		},
		{
			"Filter Test #4",
			&ffmpeg.Command{
				Duration: MustParseDuration("60.0s"),
				Width:    1920,
				Height:   1080,
			},
			"RIP",
			struct{ X, Y uint }{100, 9999},
			48,
			0x777777,
			"24.0s",
			"48.0s",
			ffmpeg.ErrInvalidFilterLocation,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := tc.command.AddDrawtextFilter(
				tc.value,
				tc.location,
				tc.size,
				tc.color,
				tc.start,
				tc.end,
			)
			assert.Equal(t, err, tc.err)
		})
	}
}
