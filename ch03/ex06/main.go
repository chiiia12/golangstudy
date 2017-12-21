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

	dx := []float64{-0.5, 0.5, -0.5, 0.5}
	dy := []float64{0.5, -0.5, 0.5, -0.5}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b uint8 = 0, 0, 0
			for i := 0; i < 4; i++ {
				y := (float64(py)+dx[i])/height*(ymax-ymin) + ymin
				x := (float64(px)+dy[i])/width*(xmax-xmin) + xmin
				z := complex(x, y)
				c := mandelbrot(z)
				sr, sg, sb, _ := c.RGBA()
				r += uint8(sr)
				g += uint8(sg)
				b += uint8(sb)

			}
			r = r / 4
			g = g / 4
			b = b / 4
			color := color.RGBA{r, g, b, 0xff}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, color)
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
