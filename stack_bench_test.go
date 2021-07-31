package benchmark

import "testing"

type computer = func(i int) int

func plusOne(i int) int {
	return i + 1
}

func makeStack(N int) computer {
	stack := plusOne

	for i := 0; i < N; i++ {
		last := stack
		stack = func(i int) int {
			return last(i + 1)
		}
	}

	return stack
}

func makeStack2(N int) computer {
	branch := func(wrapped computer) computer {
		return func(i int) int {
			return wrapped(plusOne(i))
		}
	}

	result := plusOne
	for i := 0; i < N; i++ {
		result = branch(result)
	}

	return result
}

func makeLoop(N int) computer {
	return func(j int) int {
		for i := 0; i < N; i++ {
			j = plusOne(j)
		}
		return j
	}
}

func BenchmarkStack(b *testing.B) {
	n := 100

	stack := makeStack(n)
	stack2 := makeStack2(n)
	loop := makeLoop(n)

	b.Run("A:loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			loop(i)
		}
	})
	b.Run("B:stack", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stack(i)
		}
	})
	b.Run("C:stack2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stack2(i)
		}
	})
}
