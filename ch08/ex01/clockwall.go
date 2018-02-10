package main

import (
	"net"
	"log"
	"os"
	"io"
	"fmt"
	"strings"
)

type PlaceInfo struct {
	Place   string
	Address string
}

func main() {
	fmt.Println(os.Args[1:])
	argList := parsePlace(os.Args[1:])
	fmt.Printf("argList is %v\n", argList)
	for _, v := range argList {
		conn, err := net.Dial("tcp", v.Address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		mustCopy(os.Stdout, conn)
	}
}
func parsePlace(s []string) []*PlaceInfo {
	var argList []*PlaceInfo
	for _, v := range s {
		i := strings.Index(v, "=")
		if i != -1 {
			placeInfo := PlaceInfo{v[:i], v[i+1:]}
			argList = append(argList, &placeInfo)
		}
	}
	return argList
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
