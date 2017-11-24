package main

import (
	"fmt"
	"image"
	"sync"
	"image/color"
	"net/http"
	"log"
	"io"
	"image/gif"
	"math"
	"math/rand"
	"strconv"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.Black, color.RGBA{0xFF, 0x5E, 0x19, 0xff}, color.RGBA{0x28, 0xAF, 0x78, 0xff}, color.RGBA{0x1C, 0x05, 0xFF, 0xff}, color.RGBA{0xFF, 0xFF, 0x00, 0xff}, color.RGBA{0x7F, 0x4C, 0x72, 0xff}}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query()
		fmt.Println(`queryParam.get("cycles")`, queryParam.Get("cycles"))
		var c, _ = strconv.ParseFloat(queryParam.Get("cycles"), 64)
		var res, _ = strconv.ParseFloat(queryParam.Get("res"), 64)
		var size, _ = strconv.Atoi(queryParam.Get("size"))
		var nframes, _ = strconv.Atoi(queryParam.Get("nframes"))
		var delay, _ = strconv.Atoi(queryParam.Get("delay"))

		lissajous1(w, c, res, size, nframes, delay)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous1(out io.Writer, c float64,r float64, s int,n int,d int) {
	//const (
	//	res     = 0.001
	//	size    = 100
	//	nframes = 64
	//	delay   = 8
	//)

	var cycles =c
	var res = r
	var size = s
	var  nframes =n
	var delay = d

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(math.Mod(t, 6))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
