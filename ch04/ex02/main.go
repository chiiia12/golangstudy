package main

import (
	"os"
	"crypto/sha256"
	"fmt"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		printSha256(os.Args[i])
	}

}
func printSha256(s string) {
	fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
}
