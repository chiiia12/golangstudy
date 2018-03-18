package popcount

func init() {
}

func PopCount2_5(x uint64) int {
	var count = 0
	for ; x > 0; {
		x = x & (x - 1)
		count++
	}
	return count
}
