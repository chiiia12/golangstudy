package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		fmt.Println("go func1")
		for x := 0; ; x++ {
			fmt.Println("naturals has sended", x)
			naturals <- x
		}
	}()
	go func() {
		fmt.Println("go func2")
		for {
			x := <-naturals
			fmt.Println("naturals has received", x)
			squares <- x * x
		}
	}()
	//squaresを受信しないと上２つのgoroutinesが動かない
	<-squares
	//for {
	//	fmt.Println(<-squares)
	//}
}
