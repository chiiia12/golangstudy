package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
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

type ConnectionManager struct {
	conn net.Conn
	in   chan string
	out  chan string
	ack  chan struct{}
	done chan struct{}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	cm := ConnectionManager{
		conn: conn,
		in:   make(chan string),
		out:  make(chan string),
		ack:  make(chan struct{}),
		done: make(chan struct{}),
	}
	cm.Init()

	cm.Login()

	cm.out <- "sample output"
	log.Println("out sended")
	<-cm.ack
	log.Println("ack")
	log.Println(<-cm.in)
	log.Println("cm.in")
	log.Println(<-cm.in)
	log.Println("cm.in")

	log.Println(<-cm.in)
	log.Println("cm.in")

	<-cm.done
	log.Println("done")
}

const (
	ReadyForUser = 220
	NeedPassword = 331
	UserLoggedIn = 230
	NotLoggedIn  = 530
	SyntaxError  = 500
)

func (cm *ConnectionManager) Init() {
	go func() {
		defer close(cm.done)
		for {
			select {
			case mes := <-cm.out:
				log.Println("out has received")
				fmt.Fprintf(cm.conn, mes)
				cm.ack <- struct{}{}
			case <-cm.done:
				return
			}
		}
	}()
	go func() {
		defer cm.conn.Close()
		input := bufio.NewScanner(cm.conn)
		for input.Scan() {
			cm.in <- input.Text()
			log.Println("in has inputted")
		}
	}()

}
func (cm *ConnectionManager) Login() {
	cm.Send(ReadyForUser, "Service ready for new user")
	inputUserName := <-cm.in
	log.Println("input is ", inputUserName)
	if inputUserName != UserName {
		cm.Send(SyntaxError, "syntax error")
	}
	cm.Send(NeedPassword, "User name okay, need password.")
	inputPassword := <-cm.in
	log.Println(inputPassword)
	if inputPassword != Password {
		cm.Send(SyntaxError, "syntax error")
	}
	cm.Send(UserLoggedIn, "User logged in, proceed.")
}
func (cm *ConnectionManager) Send(statusCode int, msg string) {
	fmt.Fprintf(cm.conn, "%d %s\n", statusCode, msg)
}

const (
	UserName = "USER user"
	Password = "PASS 12345"
)
