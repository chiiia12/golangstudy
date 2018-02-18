package main

import (
	"testing"
	"runtime"
	"io/ioutil"
	"fmt"
)

func BenchmarkGOMAXPROCS1(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}

func BenchmarkGOMAXPROCS2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}

func BenchmarkGOMAXPROCS3(b *testing.B) {
	runtime.GOMAXPROCS(3)
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}
func BenchmarkGOMAXPROCS5(b *testing.B) {
	runtime.GOMAXPROCS(5)
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}
func BenchmarkGOMAXPROCS6(b *testing.B) {
	runtime.GOMAXPROCS(6)
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}
func BenchmarkGOMAXPROCS10(b *testing.B) {
	runtime.GOMAXPROCS(10)
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}
func BenchmarkGOMAXPROCSNumCPU(b *testing.B) {
	fmt.Printf("runtime.NumCPU is %v\n", runtime.NumCPU()) //4
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		run(ioutil.Discard)
	}
}


//BenchmarkGOMAXPROCS1-4                 1        1310585454 ns/op
//BenchmarkGOMAXPROCS2-4                 2         670684814 ns/op
//testing: BenchmarkGOMAXPROCS2-4 left GOMAXPROCS set to 2
//BenchmarkGOMAXPROCS3-4                 3         504608177 ns/op
//testing: BenchmarkGOMAXPROCS3-4 left GOMAXPROCS set to 3
//BenchmarkGOMAXPROCS5-4                 3         393870309 ns/op
//testing: BenchmarkGOMAXPROCS5-4 left GOMAXPROCS set to 5
//BenchmarkGOMAXPROCS6-4                 3         383167047 ns/op
//testing: BenchmarkGOMAXPROCS6-4 left GOMAXPROCS set to 6
//BenchmarkGOMAXPROCS10-4                3         380104861 ns/op
//testing: BenchmarkGOMAXPROCS10-4 left GOMAXPROCS set to 10
//BenchmarkGOMAXPROCSNumCPU-4 			 3         407612502 ns/op

