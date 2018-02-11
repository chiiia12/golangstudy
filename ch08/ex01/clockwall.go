package main

import (
	"net"
	"log"
	"os"
	"fmt"
	"strings"
	"bufio"
)

type PlaceInfo struct {
	Place   string
	Address string
}
type time struct {
	index int
	time  string
}

func main() {
	out := make(chan time, 3)
	fmt.Println(os.Args[1:])
	argList := parsePlace(os.Args[1:])
	fmt.Printf("argList is %v\n", argList)
	for i, v := range argList {
		go getServerOutput(v, out, i)
	}
	for {
		times := make([]string, len(argList))
		for i := 0; i < len(argList); i++ {
			t := <-out
			times[t.index] = t.time
		}
		fmt.Println(strings.Join(times, " "))
	}
}
func getServerOutput(v *PlaceInfo, out chan time, i int) {
	conn, err := net.Dial("tcp", v.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		out <- time{i, fmt.Sprintf("%s: %s ", v.Place, string(bytes))}

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
