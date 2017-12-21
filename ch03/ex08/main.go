package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"math/big"
	"math/cmplx"
	"flag"
)

type ComplexFloat struct {
	re big.Float
	im big.Float
}

type ComplexRat struct {
	re big.Rat
	im big.Rat
}

var data = flag.String("data", "complex64", "select data type :complex64/complex128/big.Float/big.Rat")

func main() {
	flag.Parse()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.

			switch *data {
			case "complex64":
				//complex64
				zcomplex := complex(x, y)
				img.Set(px, py, mandelbrotComplex64(complex64(zcomplex)))
			case "complex128":
				//complex128
				zcomplex := complex(x, y)
				img.Set(px, py, mandelbrotComplex128(zcomplex))
			case "big.Float":
				///big.Float
				xFloat := new(big.Float)
				xFloat.SetFloat64(x)
				yFloat := new(big.Float)
				yFloat.SetFloat64(y)
				zfloat := ComplexFloat{re: *xFloat, im: *yFloat}
				img.Set(px, py, mandelbrotFloat(zfloat))
			case "big.Rat":
				//big.Rat
				xRat := new(big.Rat)
				xRat.SetFloat64(x)
				yRat := new(big.Rat)
				yRat.SetFloat64(y)
				zRat := ComplexRat{re: *xRat, im: *yRat}
				img.Set(px, py, mandelbrotRat(zRat))
			}
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotFloat(z ComplexFloat) color.Color {
	const iterations = 64
	const contrast = 15
	const absVal = 2
	var absValSq big.Float
	absValSq.SetInt64(absVal * absVal)
	var v ComplexFloat
	for n := uint8(0); n < iterations; n++ {
		var aa, ab, bb, res big.Float
		//v = v*v + z
		//v*v
		//(a+bi)(a+bi)=a^2+2abi-b^2
		aa.Mul(&v.re, &v.re)
		ab.Mul(&v.re, &v.im)
		bb.Mul(&v.im, &v.im)
		//v*v+z
		//(a+bi)(a+bi)+(c+di)=(a^2+b^2+c)+(2ab+c)i
		v.re.Sub(&aa, &bb)
		v.im.Add(&ab, &ab)

		v.re.Add(&v.re, &z.re)
		v.im.Add(&v.im, &z.im)

		//絶対値を出す
		var a, b big.Float
		a.Mul(&v.re, &v.re)
		b.Mul(&v.im, &v.im)
		res.Add(&a, &b)

		if res.Cmp(&absValSq) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotRat(z ComplexRat) color.Color {
	const iterations = 64
	const contrast = 15
	const absVal = 2
	var absValSq big.Rat
	absValSq.SetInt64(absVal * absVal)
	var v ComplexRat
	for n := uint8(0); n < iterations; n++ {
		var aa, ab, bb, res big.Rat
		//v = v*v + z
		//v*v
		//(a+bi)(a+bi)=a^2+2abi-b^2
		aa.Mul(&v.re, &v.re)
		ab.Mul(&v.re, &v.im)
		bb.Mul(&v.im, &v.im)
		//v*v+z
		//(a+bi)(a+bi)+(c+di)=(a^2+b^2+c)+(2ab+c)i
		v.re.Sub(&aa, &bb)
		v.im.Add(&ab, &ab)

		v.re.Add(&v.re, &z.re)
		v.im.Add(&v.im, &z.im)

		//絶対値を出す
		var a, b big.Rat
		a.Mul(&v.re, &v.re)
		b.Mul(&v.im, &v.im)
		res.Add(&a, &b)

		if res.Cmp(&absValSq) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
func mandelbrotComplex64(z complex64) color.Color {

	const iterations = 64
	const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		re, im := real(v), imag(v)
		if re*re+im*im > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
func mandelbrotComplex128(z complex128) color.Color {

	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
