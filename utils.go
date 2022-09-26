package stringFormatter

import "math"

type Numeric interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint
}

const maxNumericLength = 20

func NumericToStr[T Numeric](value T, signed bool, base T) string {
	// we make 20 symbols because uint64 max is 18446744073709551615  = 20 digits
	actual := value
	str := make([]byte, maxNumericLength)
	neg := false
	digits := 0
	if signed {
		if value < 0 {
			actual = T(math.Abs(float64(actual)))
			neg = true
		}
	}

	for i, _ := range str {
		digit := actual % base
		actual /= base
		if digit <= 9 {
			str[maxNumericLength-i-1] = '0' + byte(digit)
		} else {
			str[maxNumericLength-i-1] = 'A' + byte(digit) - 10
		}
		digits++
		if actual == 0 {
			break
		}
	}

	if neg {
		str[maxNumericLength-digits-1] = '-'
		digits++
	}

	return string(str)[maxNumericLength-digits:]
}
