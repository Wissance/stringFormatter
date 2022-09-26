package stringFormatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericToString(t *testing.T) {
	convertUnsignedAndCheck(t, 9, 10, "9")
	convertUnsignedAndCheck(t, 1234567890, 10, "1234567890")
	convertUnsignedAndCheck(t, 0xF1112DCB, 16, "F1112DCB")
	convertUnsignedAndCheck(t, 0o572643, 8, "572643")

	convertSignedAndCheck(t, -1234, 10, "-1234")
	convertSignedAndCheck(t, -0xFAC560, 16, "-FAC560")
}

func convertUnsignedAndCheck[TN Numeric](t *testing.T, number TN, base TN, expectedStr string) {
	result := NumericToStr(number, false, base)
	assert.Equal(t, expectedStr, result)
}

func convertSignedAndCheck[TN Numeric](t *testing.T, number TN, base TN, expectedStr string) {
	result := NumericToStr(number, true, base)
	assert.Equal(t, expectedStr, result)
}
