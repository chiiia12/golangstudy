package main

import (
	"net"
	"log"
	"io"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8033")
	done := make(chan struct{})
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
		go func() {
			<-done
			return
		}()
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	cm := CtrlConnectionManager{
		conn: conn,
		in:   make(chan string),
		out:  make(chan string),
		ack:  make(chan struct{}),
		done: make(chan struct{}),
		dir:  "/",
		dataConn: DataConnectionManager{
			out:  make(chan io.Reader),
			ack:  make(chan struct{}),
			done: make(chan struct{}),
		},
	}
	cm.Init()
	cm.Login()
	cm.HandleCommand()
}

