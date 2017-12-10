package main

import "fmt"

const(
	_ = 1<<(3*iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main(){
	fmt.Printf("%b\n",KB)
	fmt.Printf("%b\n",MB)
	fmt.Printf("%b\n",GB)
	fmt.Printf("%b\n",TB)
	fmt.Printf("%b\n",PB)
	fmt.Printf("%b\n",EB)
	fmt.Printf("%b\n",ZB)
	fmt.Printf("%b\n",YB)
}
