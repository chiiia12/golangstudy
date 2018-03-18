package popcount

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i%1)
	}
}
func PopCount2_4(x uint64) int {
	var tmp byte = 0
	for i := 0; i < 64; i++ {
		tmp += pc[byte(x>>uint64(i))]
	}
	return int(tmp)
}
