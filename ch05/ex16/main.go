package main

func join(sep string, vals ...string) string {
	//strings.Join(vals,sep)でもいけるよ
	str := ""
	for i, v := range vals {
		str += v
		if i+1 != len(vals) {
			str += sep
		}
	}
	return str
}
