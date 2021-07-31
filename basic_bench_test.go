package benchmark

import "testing"

func BenchmarkBasic(b *testing.B) {
	f := func(i int) int { return i + 1 }

	b.Run("plus", func(b *testing.B) {
		j := 0
		for i := 0; i < b.N; i++ {
			j++
		}
		_ = j
	})

	b.Run("call func", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = f(i)
		}
	})
}

