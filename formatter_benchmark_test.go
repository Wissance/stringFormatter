package stringFormatter

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkFormat3Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format("Today is : {0}, atmosphere pressure is : {1} mmHg, temperature: {2}, location: {3}", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFormatComplex3Arg(b *testing.B) {
	args := map[string]interface{}{"temperature": -10, "location": "Yekaterinburg", "time": time.Now().String(), "pressure": 725, "humidity": 34}
	for i := 0; i < b.N; i++ {
		_ = FormatComplex("Today is : {time}, atmosphere pressure is : {pressure} mmHg, temperature: {temperature}, location: {location}", args)
	}
}

func BenchmarkFmt3Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today is : %s, atmosphere pressure is : %d mmHg, temperature: %f, location: %s", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFormat5Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format("Today is : {0}, atmosphere pressure is : {1} mmHg, temperature: {2}, location: {3}, coord:{4}-{5}",
			time.Now().String(), 725, -1.54, "Yekaterinburg", "64.245", "37.895")
	}
}

func BenchmarkFmt5Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today is : %s, atmosphere pressure is : %d mmHg, temperature: %f, location: %s, coords: %s-%s",
			time.Now().String(), 725, -1.54, "Yekaterinburg", "64.245", "37.895")
	}
}
