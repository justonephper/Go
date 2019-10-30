package benchmark_test

import (
	"bytes"
	"testing"
)

func BenchmarkConcatStringByAdd(b *testing.B) {
	names := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ret := ""
		for _, val := range names {
			ret += val
		}
	}

	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	names := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range names {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
