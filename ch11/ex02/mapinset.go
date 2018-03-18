package main

import (
	"strings"
	"sort"
	"fmt"
)

type MapIntSet struct {
	set map[int]bool
}

func (m *MapIntSet) Has(x int) bool {
	_, ok := m.set[x]
	return ok
}

func (m *MapIntSet) Add(x int) {
	if m.set == nil {
		m.set = make(map[int]bool)
	}
	m.set[x] = true
}

func (m *MapIntSet) UnionWith(t *MapIntSet) {
	for k, _ := range t.set {
		m.set[k] = true
	}
}
func (m *MapIntSet) String() string {
	var str string
	keys := []int{}
	sort.Ints(keys)

	for k, _ := range m.set {
		keys = append(keys, k)
	}
	str += strings.Trim(strings.Join(strings.Fields(fmt.Sprint(keys)), " "), "{}")
	return str
}
