package main

import "fmt"

func main() {
	fmt.Println(sample())
}
func sample() (num int) {
	defer func() {
		if err := recover(); err != nil {
			num = 5
		}
	}()
	panic("panic")
	//↑deferより前に置くとコンパイルエラーになる
}
