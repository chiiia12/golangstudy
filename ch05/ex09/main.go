package main

import "strings"

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
