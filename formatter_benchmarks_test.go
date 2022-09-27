package stringFormatter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFormatMixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format("Today is : {0}, atmosphere pressure is : {1} mmHg, temperature: {2}, location: {3}", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFormatNumerics(b *testing.B) {
	rndSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rndSource)
	pressure := rnd.Intn(760)
	temperature := rnd.Intn(30)
	humidity := rnd.Intn(100)

	for i := 0; i < b.N; i++ {
		_ = Format("Today atmosphere pressure is : {0} mmHg, temperature: {1} C, humidity: {2} %", pressure, temperature, humidity)
	}
}

func BenchmarkFmtMixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today is : %s, atmosphere pressure is : %d mmHg, temperature: %f, location: %s", time.Now().String(), 725, -1.54, "Yekaterinburg")
	}
}

func BenchmarkFmtNumerics(b *testing.B) {
	rndSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rndSource)
	pressure := rnd.Intn(760)
	temperature := rnd.Intn(30)
	humidity := rnd.Intn(100)

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Today atmosphere pressure is : %d mmHg, temperature: %d C, humidity: %d percents", pressure, temperature, humidity)
	}
}
