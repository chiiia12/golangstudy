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

var wg sync.WaitGroup

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
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second)
	}
	//Addするより前に置くとcountが0なのですぐcloseしてしまう
	go func() {
		fmt.Fprintln(c, "before wg.Wait")
		wg.Wait()
		fmt.Fprintln(c, "after wg.Wait")
		tcpConn, ok := c.(*net.TCPConn)
		if !ok {
			log.Fatal("this is not tcp connection")
		}
		tcpConn.CloseWrite()
	}()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
