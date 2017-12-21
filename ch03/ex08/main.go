package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"math/big"
)

type ComplexFloat struct {
	re big.Float
	im big.Float
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := new(big.Float)
		y.SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := new(big.Float)
			x.SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			z := ComplexFloat{re: *x, im: *y}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z ComplexFloat) color.Color {
	const iterations = 200
	const contrast = 15
	const absVal = 2
	var absValSq big.Float
	absValSq.SetInt64(absVal * absVal)
	var v ComplexFloat
	for n := uint8(0); n < iterations; n++ {
		var aa, ab, bb, abab, c, d, res big.Float
		//v = v*v + z
		//v*v
		//(a+bi)(a+bi)=a^2+2abi-b^2
		aa.Mul(&v.re, &v.re)
		ab.Mul(&v.re, &v.im)
		bb.Mul(&v.im, &v.im)
		//v*v+z
		//(a+bi)(a+bi)+(c+di)=(a^2+b^2+c)+(2ab+c)i
		c = z.re
		c.Add(&aa, &bb)

		abab.Sub(&ab, &ab)
		d.Add(&z.im, &abab)
		res.Add(&aa, &bb)

		v.re = c
		v.im = d

		if res.Cmp(&absValSq) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.Black
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{0, 255 - contrast*i, 0, 0xff}

		}
	}
	//return color.RGBA{255, 255, 255, 0}
	return color.Black
}
