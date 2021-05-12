package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)



const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	} else {
		lissajous(os.Stdout)
	}
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	var palette = make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	for i := 0; i < nframes; i++ {
		scale := float64(i)
		c := color.RGBA{uint8(55 + 200 * scale), uint8(55 + 200 * scale), uint8(55 + 200 * scale), 255}
		palette = append(palette, c)
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i%len(palette)-1)+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
