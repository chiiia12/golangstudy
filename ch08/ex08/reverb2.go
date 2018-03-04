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
	//ここが10秒たった後も生きてる
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
	alive := make(chan string, 1)
	input := bufio.NewScanner(c)
	go watch(c, alive)
	for input.Scan() {
		wg.Add(1)
		alive <- input.Text()
		go echo(c, input.Text(), 1*time.Second)
	}
	//Addするより前に置くとcountが0なのですぐcloseしてしまう
	go func() {
		fmt.Fprintln(c, "before wg.Wait")
		wg.Wait()
		fmt.Fprintln(c, "after wg.Wait")
		//tcpConn, ok := c.(*net.TCPConn)
		//if !ok {
		//	log.Fatal("this is not tcp connection")
		//}
		//tcpConn.CloseWrite()
		c.Close()
	}()
}
func watch(c net.Conn, alive chan string) {
	for {
		select {
		case <-time.After(10 * time.Second):
			log.Print("10 seconds has passed")
			c.Close()
			return
		case _, ok := <-alive:
			log.Print("alive is ", ok)
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
