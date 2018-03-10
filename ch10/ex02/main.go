package main

import (
	"fmt"
	"./driver"
	_ "./zip"
	_ "./tar"
	"flag"
)

const OUTPUT_DIR = "./out"

var (
	filetype = flag.String("type", "zip", "input filetype. default is zip")
	filename = flag.String("filename", "", "input filename.")
)

func main() {
	flag.Parse()
	unarchiver, err := driver.OpenUnArchiver(*filetype)
	if err != nil {
		fmt.Println("OpenUnArchiver has error.", err)
	}
	unarchiver.UnArchive(*filename, OUTPUT_DIR)
}
