package main

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
}
func rotate(s []int, d int) []int {
	var l = len(s)
	var tmp = 0

	for i := 0; i < d; i++ {
		tmp = s[0]
		for j := 0; j < l-1; j++ {
			s[j] = s[j+1]
		}
		s[l-1] = tmp
	}
	return s
}

