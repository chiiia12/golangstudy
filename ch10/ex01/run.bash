#!/bin/bash
go build jpeg.go
go run mandelbrot.go | ./jpeg> mandelbrot.jpg
