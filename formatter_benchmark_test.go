package stringFormatter

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format("Today is : {0}, atmosphere pressure is : {1} mmHg, temperature: {2}, location: {3}", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today is : %s, atmosphere pressure is : %d mmHg, temperature: %f, location: %s", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}
