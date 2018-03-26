package main

import (
	"testing"
	"time"
	"math/rand"
)

//var だけで中身もかける条件なんだっけ。
var seed = time.Now().UTC().UnixNano()

//こうも書ける
//func init() {
//	seed = time.Now().UTC().UnixNano()
//	log.Println("seed is initialized.", seed)
//}

func BenchmarkIntSet_Add(b *testing.B) {
	var x IntSet
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		num := rng.Intn(0x1000)
		x.Add(num)
	}
}
func BenchmarkIntSet_UnionWith(b *testing.B) {
	var x, y IntSet
	rng := rand.New(rand.NewSource(seed))
	num1 := rng.Intn(0x1000)
	num2 := rng.Intn(0x1000)
	x.Add(num1)
	y.Add(num2)
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}
func BenchmarkIntSet_Has(b *testing.B) {
	var x IntSet
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		num := rng.Intn(0x1000)
		x.Has(num)
	}
}
func BenchmarkMapSet_Add(b *testing.B) {
	var x MapIntSet
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		num := rng.Intn(0x1000)
		x.Add(num)
	}
}
func BenchmarkMapSet_UnionWith(b *testing.B) {
	var x, y MapIntSet
	rng := rand.New(rand.NewSource(seed))
	num1 := rng.Intn(0x1000)
	num2 := rng.Intn(0x1000)
	x.Add(num1)
	y.Add(num2)
	for i := 0; i < b.N; i++ {
		x.UnionWith(&y)
	}
}
func BenchmarkMapSet_Has(b *testing.B) {
	var x MapIntSet
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		num := rng.Intn(0x1000)
		x.Has(num)
	}

}
