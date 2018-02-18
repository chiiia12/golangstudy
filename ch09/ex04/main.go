package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	msg := make(chan struct{})
	var mu sync.Mutex
	groutineCount := 0
	beforetime := time.Now()
	go func() {
		msg <- struct{}{}
	}()

	for {
		go func() {
			<-msg
			mu.Lock()
			groutineCount++
			mu.Unlock()

			fmt.Printf("goroutineCount is %v costtime is %s\n", groutineCount, time.Since(beforetime))

			mu.Lock()
			beforetime = time.Now()
			mu.Unlock()
			msg <- struct{}{}
		}()
	}
}

//観測したgoroutine数は37000くらい
//===OUTPUT EXAMPLE===
//goroutineCount is 491 costtime is 1.394µs
//goroutineCount is 492 costtime is 4.13µs
//goroutineCount is 493 costtime is 1.418µs
//goroutineCount is 494 costtime is 7.522µs
//goroutineCount is 495 costtime is 1.717µs

