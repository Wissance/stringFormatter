package stringFormatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericToString(t *testing.T) {
	number := 1234567890
	result := numericToStr(number, false, 10)
	assert.Equal(t, "1234567890", result)
}

func convertAndTest[TN numeric](t *testing.T, )
