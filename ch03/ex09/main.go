package main

import (
	"net/http"
	"log"
	"image/color"
	"math/cmplx"
	"image"
	"image/png"
	"strconv"
)

//http://localhost:8000/?x=-3&y=-3&mag=2
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		minx, _ := strconv.ParseFloat(r.Form["x"][0], 64)
		miny, _ := strconv.ParseFloat(r.Form["y"][0], 64)
		mag, _ := strconv.Atoi(r.Form["mag"][0])
		const (
			xmax, ymax    = +2, +2
			width, height = 1024, 1024
		)

		mWidth :=mag*width
		mHeight :=mag*height
		img := image.NewRGBA(image.Rect(0, 0, mWidth, mHeight))
		//以下のheightにmagがかかってないからキャンバスだけ広げて中身がない
		for py := 0; py < mHeight; py++ {
			y := float64(py/mHeight)*(ymax-miny) + miny
			for px := 0; px < mWidth; px++ {
				x := float64(px/mWidth)*(xmax-minx) + minx
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(w, img) // NOTE: ignoring errors
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - contrast*n, 255 - contrast*n*2, 255 - contrast*n*2}
		}
	}
	//return color.YCbCr{255 - contrast,255-contrast,255-contrast}
	return color.Black
}
