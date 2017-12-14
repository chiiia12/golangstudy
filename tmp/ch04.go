package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%p\n",s)//0xc420016210
	fmt.Printf("%p\n",s[:2])//0xc420016210
	fmt.Printf("%p\n",&s[1])
	fmt.Printf("%p\n",s[1:2])
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
