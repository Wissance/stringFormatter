package stringFormatter

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkFormat4Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format("Today is : {0}, atmosphere pressure is : {1} mmHg, temperature: {2}, location: {3}", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFmt4Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today is : %s, atmosphere pressure is : %d mmHg, temperature: %f, location: %s", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFormat6Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format("Today is : {0}, atmosphere pressure is : {1} mmHg, temperature: {2}, location: {3}, coord:{4}-{5}",
			time.Now().String(), 725, -1.54, "Yekaterinburg", "64.245", "37.895")
	}
}

func BenchmarkFmt6Arg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today is : %s, atmosphere pressure is : %d mmHg, temperature: %f, location: %s, coords: %s-%s",
			time.Now().String(), 725, -1.54, "Yekaterinburg", "64.245", "37.895")
	}
}

func BenchmarkFormatComplex7Arg(b *testing.B) {
	args := map[string]interface{}{"temperature": -10, "location": "Yekaterinburg", "time": time.Now().String(), "pressure": 725, "humidity": 34,
		"longitude": "64.245", "latitude": "35.489"}
	for i := 0; i < b.N; i++ {
		_ = FormatComplex("Today is : {time}, atmosphere pressure is : {pressure} mmHg, humidity: {humidity}, temperature: {temperature}, location: {location}, coords:{longitude}-{latitude}", args)
	}
}
