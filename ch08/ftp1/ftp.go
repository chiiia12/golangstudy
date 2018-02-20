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
		log.Println("for")
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
type ConnectionManager struct{

}
func handleConn(conn net.Conn) {
	in := make(chan string)
	out := make(chan string)
	ack := make(chan struct{})
	done := make(chan struct{})
	fmt.Fprintln(conn, "handleConn")
	//go clientWriter(conn, out, ack)

	go func(done chan struct{}, conn net.Conn, in chan string) {
		defer close(done)
		for {
			select {
			case mes := <-out:
				log.Println("out has received")
				fmt.Fprintf(conn, mes)
				ack <- struct{}{}
			case <-done:
				return
			}
		}
	}(done, conn, in)
	go func(conn net.Conn, in chan string, ack chan struct{}) {
		defer conn.Close()
		input := bufio.NewScanner(conn)
		for input.Scan() {
			in <- input.Text()
			log.Println("in has inputted")
		}
	}(conn, in, ack)
	out <- "sample output"
	log.Println("out sended")
	<-ack
	log.Println("ack")
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
