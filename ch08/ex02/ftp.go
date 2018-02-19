package main

import (
	"net"
	"fmt"
	"log"
	"bufio"
	"strings"
)

type FTPServer struct {
	host string
	port int
	done chan struct{}
}

func main() {
	s := &FTPServer{
		host: "localhost",
		port: 8021,
		done: make(chan struct{}),
	}

	s.Run()
}
func (s *FTPServer) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
		go func() {
			<-s.done
			return
		}()
	}
}
func handleConn(c net.Conn) {
	defer c.Close()

	log.Println("connected")
	done := make(chan struct{})
	s := NewSession(c, done)
	s.Login()
}

type CtrlConnManager struct {
	conn net.Conn
	in   chan string
	ack  chan struct{}
	out  chan string
	done chan struct{}
}

type Session struct {
	ctrl *CtrlConnManager
	done chan struct{}
}

func (s *Session) Close() {
	close(s.done)
}

const (
	ReadyForUser = 220
	NeedPassword = 331
	UserLoggedIn = 230
)

var users = map[string]string{
	"user": "12345",
}

func (s *Session) Login() {
	s.ctrl.SendMessage(ReadyForUser, "my go ftp server ready")
	userseq := strings.Split(<-s.ctrl.out, " ")
	if !strings.EqualFold(userseq[0], "USER") {
		log.Println("invalid user name")
		return
	}

	pass := users[userseq[1]]
	s.ctrl.SendMessage(NeedPassword, "Send Password")

	passseq := strings.Split(<-s.ctrl.out, " ")
	if !strings.EqualFold(passseq[0], "PASS") {
		log.Println("password error")
		return
	}

	if pass != passseq[1] {
		return
	}
	s.ctrl.SendMessage(UserLoggedIn, "LoginSuccessful")
}

func NewSession(conn net.Conn, done chan struct{}) *Session {
	ctrl := NewCtrlConnManager(conn)
	ctrl.Run()
	s := &Session{
		ctrl: ctrl,
		done: done,
	}
	go func() {
		<-ctrl.done
		s.Close()
	}()
	return s
}
func NewCtrlConnManager(conn net.Conn) *CtrlConnManager {
	return &CtrlConnManager{
		conn: conn,
		in:   make(chan string),
		ack:  make(chan struct{}),
		out:  make(chan string),
		done: make(chan struct{}),
	}
}
func (c *CtrlConnManager) Run() {
	go func() {
		defer c.Close()
		input := bufio.NewScanner(c.conn)
		for input.Scan() {
			c.out <- input.Text()
		}
		log.Println("Goint to closing")
	}()
	go func() {
		defer c.conn.Close()
		for {
			select {
			case mes := <-c.in:
				fmt.Fprintf(c.conn, mes)
				c.ack <- struct{}{}
			case <-c.done:
				return
			}
		}
	}()
}
func (c *CtrlConnManager) Close() {
	close(c.done)
}
func (c *CtrlConnManager) SendMessage(code int, mes string) {
	c.Send(fmt.Sprintf("%d %s\n", code, mes))
}
func (c *CtrlConnManager) Send(mes string) {
	c.in <- mes
	<-c.ack
}
