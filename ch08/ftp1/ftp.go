package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8022")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	in := make(chan string)
	out := make(chan string)
	ack := make(chan struct{})
	fmt.Fprintln(conn, "handleConn")
	go clientWriter(conn, out, ack)

	input := bufio.NewScanner(conn)
	out <- "hoge"
	for input.Scan() {
		in <- input.Text()
		out <- input.Text()
		<-ack
	}
}
func clientWriter(conn net.Conn, out <-chan string, ack chan struct{}) {
	//for {
	//	select {
	//	case msg := <-out:
	//		fmt.Fprintf(conn, msg)
	//		log.Println("out")
	//		ack <- struct{}{}
	//	}
	//}
	for msg := range out {
		fmt.Fprintln(conn, msg)
		ack <- struct{}{}
	}
}
