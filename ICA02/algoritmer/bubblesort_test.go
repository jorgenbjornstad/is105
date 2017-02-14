package algoritmer

import (
	"math/rand"
	"testing"
	"time"
)

// https://golang.org/doc/effective_go.html#init
func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

// Skriv "benchmark"-tester for benchmarkBSortModified funksjonen
// Skriv en ny testfunksjon benchmarkBSortModified

func BenchmarkBSortModified100(b *testing.B) {
	benchmarkBSortModified(100, b)
}

func BenchmarkBSortModified1000(b *testing.B) {
	benchmarkBSortModified(1000, b)
}

func BenchmarkBSortModified10000(b *testing.B) {
	benchmarkBSortModified(10000, b)
}

func benchmarkBSortModified(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort_modified()
	}
}
