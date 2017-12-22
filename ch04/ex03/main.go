package main

import "fmt"

func main() {
	s := [6]int{0, 1, 2, 3, 4, 5}
	s = reverse(&s)
	fmt.Println(s)
}

func reverse(s *[6]int) [6]int {
	var arr = *s
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
