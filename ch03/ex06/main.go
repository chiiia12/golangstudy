package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			//z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, superSampling(x, y, 1))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
func superSampling(x, y, d float64) color.Color {
	r1, g1, b1, _ := mandelbrot(complex(x+d, y+d)).RGBA()
	r2, g2, b2, _ := mandelbrot(complex(x-d, y-d)).RGBA()
	r3, g3, b3, _ := mandelbrot(complex(x+d, y-d)).RGBA()
	r4, g4, b4, _ := mandelbrot(complex(x-d, y+d)).RGBA()

	return color.RGBA{uint8((r1 + r2 + r3 + r4) / 4), uint8((g1 + g2 + g3 + g4) / 4), uint8((b1 + b2 + b3 + b4) / 4), 0xff}

}
