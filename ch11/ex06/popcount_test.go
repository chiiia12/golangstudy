package popcount

import "testing"

func benchmark(b *testing.B, f func(uint64) int, x uint64) {
	for i := 0; i < b.N; i++ {
		f(x)
	}
}
func BenchmarkPopCount1(b *testing.B) {
	benchmark(b, PopCount, 0x0)

}
func BenchmarkPopCount2(b *testing.B) {
	benchmark(b, PopCount2_4, 0x0)

}
func BenchmarkPopCount3(b *testing.B) {
	benchmark(b, PopCount2_5, 0x0)

}

func BenchmarkPopCount4(b *testing.B) {
	benchmark(b, PopCount, 0xFFFF)
}
func BenchmarkPopCount5(b *testing.B) {
	benchmark(b, PopCount2_4, 0xFFFF)
}

func BenchmarkPopCount6(b *testing.B) {
	benchmark(b, PopCount2_5, 0xFFFF)
}

func BenchmarkPopCount7(b *testing.B) {
	benchmark(b, PopCount, 0xFFFFFFFFFF)
}
func BenchmarkPopCount8(b *testing.B) {
	benchmark(b, PopCount2_4, 0xFFFFFFFFFF)
}
func BenchmarkPopCount9(b *testing.B) {
	benchmark(b, PopCount2_5, 0xFFFFFFFFFF)
}
