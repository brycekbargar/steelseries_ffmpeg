package ffmpeg_test

import (
	"strings"
	"testing"
	"time"

	"github.com/brycekbargar/steelseries_ffmpeg/ffmpeg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommand_Render(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		command *ffmpeg.Command
		result  string
	}{
		{
			"Basic Command",
			&ffmpeg.Command{
				Input:    "test_input2.mp4",
				Output:   "test_output2.mp4",
				Width:    1920,
				Height:   1080,
				Duration: MustParseDuration("60.0s"),
			},
			"ffmpeg -i test_input2.mp4 test_output2.mp4",
		},
		{
			"Filter Test #1",
			&ffmpeg.Command{
				Input:    "test_input2.mp4",
				Output:   "test_output2.mp4",
				Width:    1920,
				Height:   1080,
				Duration: MustParseDuration("60.0s"),
				Filters: []ffmpeg.Filter{
					ffmpeg.DrawTextFilter{
						Value: "I’m sOoOo good at this game! xD",
						X:     200,
						Y:     100,
						Size:  64,
						Color: 0xFFFFFF,
						Start: MustParseDuration("23.0s").Seconds(),
						End:   MustParseDuration("40.0s").Seconds(),
					},
				},
			},
			"ffmpeg -i test_input2.mp4 -vf \"drawtext = enable='between(t,23.0,40.0)':text='I’m sOoOo good at this game! xD':fontcolor=0xFFFFFF:fontsize=64:x=200:y=100\" test_output2.mp4",
		},
		{
			"Filter Test #2",
			&ffmpeg.Command{
				Input:    "test_input2.mp4",
				Output:   "test_output2.mp4",
				Width:    1920,
				Height:   1080,
				Duration: MustParseDuration("60.0s"),
				Filters: []ffmpeg.Filter{
					ffmpeg.DrawTextFilter{
						Value: "Brutal, Savage, Rekt",
						X:     0,
						Y:     0,
						Size:  48,
						Color: 0x000000,
						Start: MustParseDuration("0.0s").Seconds(),
						End:   MustParseDuration("60.0s").Seconds(),
					},
				},
			},
			"ffmpeg -i test_input2.mp4 -vf \"drawtext = enable='between(t,0.0,60.0)':text='Brutal, Savage, Rekt':fontcolor=0x000000:fontsize=48:x=0:y=0\" test_output2.mp4",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			b := new(strings.Builder)
			err := tc.command.Render(b)

			require.NoError(t, err)
			assert.Equal(t, tc.result, b.String())
		})
	}
}

func MustParseDuration(s string) (d time.Duration) {
	d, _ = time.ParseDuration(s)
	return
}
