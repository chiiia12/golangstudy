package main

import "fmt"

func main(){
	s :=[]int{0,1,2,3,4,5}

	reverse(toPtr(s[:2]))
	fmt.Println(s)
	reverse(toPtr(s[2:]))
	fmt.Println(s)
	reverse(toPtr(s))
	fmt.Println(s)
}

func reverse(s *[]int) {
	var arr []int =*s
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
func toPtr(arr []int)*[]int{
	return &arr
}
