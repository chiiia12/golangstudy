package main

import (
	"fmt"
)

func main() {

	msg := make(chan struct{})
	done := make(chan struct{})
	done2 := make(chan struct{})

	go func(done chan struct{}, msg chan struct{}) {
		fmt.Println("msg has sended")
		msg <- struct{}{}
		done <- struct{}{}
	}(done, msg)

	go func(done chan struct{}, msg chan struct{}) {
		<-msg
		fmt.Println("msg has received")
		done2 <- struct{}{}
	}(done, msg)
	<-done
	<-done2
	//ココで待ち合わせしておかないとgoroutineの処理をする前にmainスレッドが終わっちゃう。
	// ココが終わっちゃうと他のゴルーチンが動いていても何も吐き出されない
}
