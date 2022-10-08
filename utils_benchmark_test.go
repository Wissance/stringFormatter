package stringFormatter

import (
	"strconv"
	"testing"
)

func BenchmarkNToS(b *testing.B) {
	for i := 1234; i < b.N*10000; i++ {
		_ = NumericToStr[int](i, 10)
	}
}

func BenchmarkStrconv(b *testing.B) {
	for i := 1234; i < b.N*10000; i++ {
		_ = strconv.FormatInt(int64(i), 10)
	}
}
