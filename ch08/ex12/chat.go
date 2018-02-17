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
	clientsName := make(map[string]bool)
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
			clientsName[cli.name] = true
			var clientsList string
			for k, v := range clientsName {
				if v {
					clientsList += k
					clientsList += " "
				}
			}
			cli.c <- "client list:" + clientsList
		case cli := <-leaving:
			clientsName[cli.name] = false
			delete(clients, cli.c)
			close(cli.c)
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
	entering <- clientset{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
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
