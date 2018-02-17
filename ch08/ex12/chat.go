package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
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

var (
	entering = make(chan client)
	leaving  = make(chan client)
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
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
func handleConn(conn net.Conn) {
	fmt.Fprintln(conn, "handleConn")
	ch := make(chan string) //送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) //ネットワークエラーを無視
	}
}
