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
