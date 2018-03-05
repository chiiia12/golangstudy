#!/bin/bash
go build main.go
go run mandelbrot.go | ./main > mandelbrot.jpg
go run mandelbrot.go | ./main -o png > mandelbrot.png
go run mandelbrot.go | ./main -o gif > mandelbrot.gif
