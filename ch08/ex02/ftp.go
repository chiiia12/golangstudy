package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"strings"
	"path/filepath"
	"strconv"
	"io"
	"bytes"
	"io/ioutil"
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

type DataConnectionManager struct {
	conn net.Conn
	out  chan io.Reader
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
		dir:  "/",
	}
	cm.Init()
	cm.Login()
	cm.HandleCommand()

}

func (cm *ConnectionManager) HandleCommand() {
	log.Println("HandleCommand")
	dataConn := &DataConnectionManager{
		out:  make(chan io.Reader),
		ack:  make(chan struct{}),
		done: make(chan struct{}),
	}
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
		case "PORT":
			arg := strings.Split(command[1], ",")
			address := strings.Join(arg[0:4], ".")
			a, _ := strconv.Atoi(arg[4])
			b, _ := strconv.Atoi(arg[5])
			port := a*256 + b
			log.Println(address)
			log.Println(port)
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
			if err != nil {
				log.Println(err)
			}
			dataConn.conn = conn
			dataConn.Init()
			log.Println("port")
			cm.Send(CommandSucceeded, "port command succeed")
		case "LIST":
			cm.Send(TransferStarting, "transfer start")
			if dataConn.conn == nil {
				log.Println("dataConn.conn is nil")
			}
			files, err := ioutil.ReadDir("/")
			if err != nil {
				log.Println(err)
			}
			for _, v := range files {
				p := fmt.Sprintf("%s\t%s\t%d\t%s\n", v.Mode(), v.ModTime(), v.Size(), v.Name())
				dataConn.out <- bytes.NewBufferString(p)
				<-dataConn.ack
			}
			//何かのデータを転送する
			log.Println("ack has received")
			close(dataConn.done)
			log.Println("send transfer complete")
			dataConn.conn.Close()
			cm.Send(ClosingConnection, "Transfer complete")
			log.Println("after send")
		case "CWD":
			cm.dir = filepath.Join(cm.dir, command[1])
			log.Println(cm.dir)
			cm.Send(CommandSucceeded, fmt.Sprintf("%s is the current directory.", cm.dir))
		default:
			cm.Send(NotImplemented, "command not implemented")
		}
	}
}
func (dm *DataConnectionManager) Init() {
	go func() {
		for {
			select {
			case mes := <-dm.out:
				io.Copy(dm.conn, mes)
				dm.ack <- struct{}{}
			case <-dm.done:
				return
			}
		}
	}()
}
func (cm *ConnectionManager) Init() {
	go func() {
		defer close(cm.done)
		for {
			select {
			case mes := <-cm.out:
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
