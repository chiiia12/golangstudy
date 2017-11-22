package popcount

func init() {
}

func PopCountEx05(x uint64) int {
	var count = 0
	for ; x > 0; {
		x = x & (x - 1)
		count++
	}
	return count
}
