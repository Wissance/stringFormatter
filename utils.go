package stringFormatter

type numeric interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint
}

const maxNumericLength = 20

func numericToStr[T numeric](value T, signed bool, base T) string {
	// we make 20 symbols because uint64 max is 18446744073709551615  = 20 digits
	actual := value
	str := make([]byte, maxNumericLength)
	neg := false
	digits := 0
	if signed {
		if value < 0 {
			actual = value
			neg = true
		}
	}

	for i, _ := range str {
		digit := actual % base
		actual /= base
		str[maxNumericLength-i-1] = 0x30 + byte(digit)
		digits++
		if actual == 0 {
			break
		}
	}

	if neg {
		str[digits+1] = '-'
		digits++
	}

	return string(str)[maxNumericLength-digits:]
}
