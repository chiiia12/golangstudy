package main

import (
	"os"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"flag"
	"image/png"
	"log"
	"image/gif"
)

var (
	output = flag.String("o", "jpeg", "option -o is output format. default is jpeg")
)

func main() {
	flag.Parse()
	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	fmt.Fprintln(os.Stderr, "output is ", *output)

	switch *output {
	case "png":
		err = png.Encode(os.Stdout, img)
	case "jpeg":
		err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95})
	case "gif":
		err = gif.Encode(os.Stdout, img, &gif.Options{NumColors: 256})
	default:
		err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 95})
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
