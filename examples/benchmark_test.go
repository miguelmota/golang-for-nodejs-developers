package example

import (
	"testing"
)

func BenchmarkFibRec(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibRec(10)
	}
}

func BenchmarkFibLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibLoop(10)
	}
}

func fibRec(n int) int {
	if n <= 1 {
		return n
	}

	return fibRec(n-1) + fibRec(n-2)
}

func fibLoop(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
