package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

type clientset struct {
	c    client
	name string
}

var (
	entering = make(chan clientset)
	leaving  = make(chan clientset)
	messages = make(chan string) //クライアントから受信するすべてのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool) //すべての接続されているクライアント
	for {
		select {
		case msg := <-messages:
			//受信するメッセージをすべてのクライアントの送信用メッセージチャネルへブロードキャストする
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli.c] = true
		case cli := <-leaving:
			delete(clients, cli.c)
			close(cli.c)
		}
	}
}
func handleConn(conn net.Conn) {
	fmt.Fprintln(conn, "handleConn")
	ch := make(chan string) //送信用のクライアントメッセージ
	ok := make(chan struct{})
	go clientWriter(conn, ch)

	entering <- clientset{ch, conn.RemoteAddr().String()}
	input := bufio.NewScanner(conn)
	ch <- "input your name:"
	var who string
	if input.Scan() {
		who = input.Text()
	}
	ch <- "You are " + who
	messages <- who + " has arrived"

	go func(conn net.Conn) {
		for {
			select {
			case <-time.After(5 * time.Minute):
				conn.Close()
				return
			case <-ok:
				//なにもしない
			}
		}
	}(conn)
	for input.Scan() {
		ok <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	leaving <- clientset{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) //ネットワークエラーを無視
	}
}
