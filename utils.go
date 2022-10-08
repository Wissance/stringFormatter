package stringFormatter

import "math"

type Numeric interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint
}

const maxNumericLength = 20

/*var digitsTable = map[int]byte{
	0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9',
	10: 'A', 11: 'B', 12: 'C', 13: 'D', 14: 'E', 15: 'F',
}*/

func NumericToStr[T Numeric](value T /*signed bool,*/, base T) string {
	// we make 20 symbols because uint64 max is 18446744073709551615  = 20 digits
	actual := value
	var str [maxNumericLength]byte
	strValueLen := len(str)
	//make([]byte, maxNumericLength)
	neg := false
	digits := 0
	if value < 0 {
		actual = T(math.Abs(float64(actual)))
		neg = true
	}

	var digit byte
	for i, _ := range str {
		digit = byte(actual % base)
		actual /= base
		// str[maxNumericLength-i-1] = digitsTable[int(digit)] // slow!!!
		if digit <= 9 {
			str[strValueLen-i-1] = '0' + digit
		} else {
			str[strValueLen-i-1] = 'A' + digit - 10
		}
		digits++
		if actual == 0 {
			break
		}

	}

	if neg {
		str[strValueLen-digits-1] = '-'
		digits++
	}

	return string(str[strValueLen-digits:])
}
