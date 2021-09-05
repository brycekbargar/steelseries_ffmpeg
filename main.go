package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/brycekbargar/steelseries_ffmpeg/ffmpeg"
)

type Resolution struct {
	X uint
	Y uint
}
type Video struct {
	Input      string
	Duration   string
	Resolution Resolution
	Output     string
}
type TextEffect struct {
	Text      string
	X         uint
	Y         uint
	FontSize  uint
	FontColor string
	Start     string
	End       string
}
type TextEffectsRequest struct {
	Video      Video
	TextEffect TextEffect
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "For Glory!")
	})
	http.HandleFunc("/api/texteffects", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var r TextEffectsRequest
		err := json.NewDecoder(req.Body).Decode(&r)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		c, err := ffmpeg.NewCommand(
			ffmpeg.FilePath(r.Video.Input),
			ffmpeg.FilePath(r.Video.Output),
			struct{ X, Y uint }{r.Video.Resolution.X, r.Video.Resolution.Y},
			r.Video.Duration,
		)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		col, err := strconv.ParseUint(r.TextEffect.FontColor, 16, 32)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		err = c.AddDrawtextFilter(
			r.TextEffect.Text,
			struct{ X, Y uint }{r.TextEffect.X, r.TextEffect.Y},
			r.TextEffect.FontSize,
			uint(col),
			r.TextEffect.Start,
			r.TextEffect.End,
		)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		b := new(strings.Builder)
		err = c.Render(b)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(b.String())
	})

	log.Fatal(http.ListenAndServe(":4123", nil))
}
