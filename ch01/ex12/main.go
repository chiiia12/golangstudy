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
	"reflect"
	"strings"
	"strconv"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.Black, color.RGBA{0xFF, 0x5E, 0x19, 0xff}, color.RGBA{0x28, 0xAF, 0x78, 0xff}, color.RGBA{0x1C, 0x05, 0xFF, 0xff}, color.RGBA{0xFF, 0xFF, 0x00, 0xff}, color.RGBA{0x7F, 0x4C, 0x72, 0xff}}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query()
		fmt.Println(`queryParam.get("cycles")`, queryParam.Get("cycles"))
		var f, _ = strconv.ParseFloat(queryParam.Get("cycles"), 64)
		fmt.Println(`f is %q`, f);
		fmt.Println(`f's type is %q`, reflect.TypeOf(f));
		lissajous1(w, f)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] =%q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "r.Form[cycles] is %q", r.Form["cycles"])
	fmt.Fprintf(w, "r.Form[cycles]'s type  is %q", reflect.TypeOf(r.Form["cycles"]))
	var f, _ = strconv.ParseFloat(strings.Join(r.Form["cycles"], ""), 64)
	lissajous(w, f)
}

func lissajous(out io.Writer, c float64) {
	const (
		cycles  = 5
		res     = 0.00
		size    = 100
		nframes = 64
		delay   = 8
	)
	//if c == 0 {
	c = cycles
	//}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < c*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(math.Mod(t, 6))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
func lissajous1(out io.Writer, f float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	var cycles = f

	fmt.Println(`lissajous1's f is %q`, f)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(math.Mod(t, 6))+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
