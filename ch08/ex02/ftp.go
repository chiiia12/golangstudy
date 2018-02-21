package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
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
	dir  string
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	cm := ConnectionManager{
		conn: conn,
		in:   make(chan string),
		out:  make(chan string),
		ack:  make(chan struct{}),
		done: make(chan struct{}),
		dir:  "/",
	}
	cm.Init()
	cm.Login()
	cm.HandleCommand()

}

func (cm *ConnectionManager) HandleCommand() {
	for {
		msg := <-cm.in
		log.Println(msg)
		command := strings.Split(msg, " ")
		switch(command[0]) {
		case "SYST":
			cm.Send(SystemType, "UNIX Type: L8")
		case "FEAT":
			cm.Send(SystemStatus, "End FEAT.")
		case "PWD":
			cm.Send(PathNameCreated, "\""+cm.dir+"\" is the current directory.")
		case "EPSV":
			cm.Send(EnteringPassive, "Entering extended passive mode")
			//つなぎにいく？
			//conn, err := net.Dial("tcp", cm.conn.RemoteAddr().String()+":")
		case "PORT":
			log.Println("port")
		case "LIST":
			cm.Send(TransferStarting, "Data connection already open.Transfer starting.")
			//files, err := ioutil.ReadDir("/")
			//if err != nil {
			//	log.Println(err)
			//}
			//for v := range files {
			//	log.Println(v)
			//}
			//何かのデータを転送する
			cm.Send(ClosingConnection, "Transfer complete")

		case "CWD":
			cm.dir = filepath.Join(cm.dir, command[1])
			log.Println(cm.dir)
			cm.Send(CommandSucceeded, fmt.Sprintf("%s is the current directory.", cm.dir))
		}
	}
}
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
	log.Println(inputUserName)
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
	log.Println("login succeeded")
}
func (cm *ConnectionManager) Send(statusCode int, msg string) {
	//fmt.Fprintf(cm.conn, "%d %s\n", statusCode, msg)
	cm.out <- fmt.Sprintf("%d %s\n", statusCode, msg)
	<-cm.ack
}
