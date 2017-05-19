package functions

import (
	"testing"
	"fmt"
)






func BenchmarkRuneScanner(b *testing.B) {

	for i := 100; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		RuneScanner()
		fmt.Println("hallo")

	}
}
