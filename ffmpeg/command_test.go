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
