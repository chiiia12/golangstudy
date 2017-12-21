package main

func main() {
	var s = []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8}
	s = removeDepulicate(s, 0, 8)
}
func removeDepulicate(s []int, first int, end int) []int {
	tmp := -1
	for i := first; i < end; i++ {
		if tmp == s[i] {
			s = remove(s, i)
			tmp = -1
			end--
		}
		tmp = s[i]
	}
	return s
}

func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func toPtr(ints []int) *[]int {
	return &ints
}