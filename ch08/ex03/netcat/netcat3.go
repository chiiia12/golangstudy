package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	//netcat3
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok{
		log.Fatal("This is not TCPConnection")
	}

	done := make(chan struct{})
	go func() {
		//connectionがcloseするまでこの先にいかない
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	//標準入力を閉じるまでここには来ない
	tcpConn.CloseWrite()
	log.Println("after mustCopy")
	//doneを受信するまでその先に進まない
	<-done
	log.Print("after done")
}

//reverb1サーバー内の内容を1回出力しているだけ
func mustCopy(dst io.Writer, src io.Reader) {
	log.Print("mustCopy")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
