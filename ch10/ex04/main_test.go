package main

import (
	"testing"
)

//推移的依存までは見てない
var hashDep = []string{
	"errors",
	"internal/race",
	"io",
	"runtime",
	"runtime/internal/atomic",
	"runtime/internal/sys",
	"sync",
	"sync/atomic",
	"unsafe",
}

func TestGetGoTest(t *testing.T) {
	var dependencyMap = make(map[string]struct{})
	getGoList("hash", &dependencyMap)
	for _, v := range hashDep {
		if _, ok := dependencyMap[v]; !ok {
			t.Errorf("dependencyMap[%v] is not ok.", v)
		}
	}
}
