//3章の内容の試しコードとメモ
package main

import "fmt"

func main() {
	var hoge int8 = 1 //符号付き整数
	var fuga uint = 1 //符号なし整数
	//uint = -1を代入するとコンパイルエラー "constant -1 overflows uint"
	fmt.Println("int8 %v", hoge)
	fmt.Println("uint8 %v", fuga)

	var hoge2 rune = 3 //int32のsynonym
	var fuga2 byte = 4 //uint8のsynonym

	fmt.Println("rune %v", hoge2)
	fmt.Println("byte %v", fuga2)

	var x uint8 = 1<<1 | 1<<5
	fmt.Println("x is %v", x)

	//var y int8 = 1<<-1//シフトするビットは符号なし整数でないといけない
	//constant -1 overflows uint
	//invalid negative shift count -1
	//fmt.Println("y is %v", y)

	var f float32 = 16777216
	fmt.Println("f is %v f+1 is %v",f,f+1)
	fmt.Println(f == f+1)
	//1.6777216e+07 1.6777216e+07
	//TODO:内部的にどうなってるのか？切り捨てされているわけではないのか？



}
