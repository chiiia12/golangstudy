package main

import (
	"crypto/sha256"
	"fmt"
	"crypto/sha512"
	"flag"
)

var (
	optSha384 = flag.Bool("a", false, "option -a use SHA384")
	optSha512 = flag.Bool("b", false, "option -b use SHA512")
	str       = flag.String("s", "", "option -s is string")
)

func main() {
	flag.Parse()

	if *optSha384 {
		printSha384(*str)
	} else if *optSha512 {
		printSha512(*str)
	} else {
		printSha256(*str)
	}

}
func printSha256(s string) {
	fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
}
func printSha384(s string) {
	fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
}
func printSha512(s string) {
	fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
}
