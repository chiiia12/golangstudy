package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {

	msg := make(chan struct{})
	done := make(chan struct{}, 2)
	end := make(chan struct{})

	countConnection := 0
	var mu sync.Mutex

	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Printf("connection count is %v\n", countConnection)
			end <- struct{}{}
		}
	}()

	go func(done chan struct{}, msg chan struct{}) {
		defer func() { done <- struct{}{} }()
		msg <- struct{}{}
		for {
			select {
			case <-end:
				fmt.Println("end")
				return
			default:
				<-msg
				msg <- struct{}{}
				mu.Lock()
				countConnection++
				mu.Unlock()
			}
		}
	}(done, msg)

	go func(done chan struct{}, msg chan struct{}) {
		defer func() { done <- struct{}{} }()
		for {
			select {
			case <-end:
				fmt.Println("end")
				return
			default:
				<-msg
				msg <- struct{}{}
				mu.Lock()
				countConnection++
				mu.Unlock()

			}
		}
	}(done, msg)

	<-done
	fmt.Println("done")
	//ココで待ち合わせしておかないとgoroutineの処理をする前にmainスレッドが終わっちゃう。
	// ココが終わっちゃうと他のゴルーチンが動いていても何も吐き出されない
}
