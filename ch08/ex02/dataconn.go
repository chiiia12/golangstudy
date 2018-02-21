package main

import (
	"io"
	"net"
)

type DataConnectionManager struct {
	conn net.Conn
	out  chan io.Reader
	ack  chan struct{}
	done chan struct{}
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
