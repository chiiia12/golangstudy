package main

import (
	"net"
	"log"
	"bufio"
	"time"
	"fmt"
	"strings"
	"sync"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	var wg sync.WaitGroup
	func() {
		wg.Add(1)
		fmt.Fprintln(c, "\t", wg.)
		defer wg.Done()
		fmt.Fprintln(c, "\t", strings.ToUpper(shout))
		time.Sleep(delay)
		fmt.Fprintln(c, "\t", shout)
		time.Sleep(delay)
		fmt.Fprintln(c, "\t", strings.ToLower(shout))
	}()
}
