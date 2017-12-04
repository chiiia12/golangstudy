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
		if err:=r.ParseForm();err !=nil{
			log.Print(err)
		}
		w.Header().Set("Content-Type","image/svg+xml")

	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func output(){
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			flg := true
			ax, ay, flg := corner(i+1, j)
			bx, by, flg := corner(i, j)
			cx, cy, flg := corner(i, j+1)
			dx, dy, flg := corner(i+1, j+1)
			if (flg) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")

}

func corner(i, j int) (float64, float64, bool) {
	isCorrectValue := true
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	if (math.IsInf(z, 0)) {
		isCorrectValue = false
	}
	return sx, sy, isCorrectValue
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
