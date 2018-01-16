package main

func join(sep string, vals ...string) string {
	str := ""
	for i, v := range vals {
		str += v
		if i+1 != len(vals) {
			str += sep
		}
	}
	return str
}

