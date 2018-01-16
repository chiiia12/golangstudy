package main

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("param size is 0")
	}
	max := vals[0]
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("param size is 0")
	}
	min := vals[0]
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
