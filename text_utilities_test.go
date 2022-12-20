package stringFormatter

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMapToString(t *testing.T) {
	options := map[string]interface{}{
		"connectTimeout": 1000,
		"useSsl":         true,
		"login":          "sa",
		"password":       "sa",
	}

	str := MapToString(&options, KeyValueWithSemicolonSepFormat, ", ")
	assert.True(t, len(str) > 0)
	// we check only parts because range produce string in random order
	assert.True(t, strings.Contains(str, "connectTimeout : 1000"))
	assert.True(t, strings.Contains(str, "useSsl : true"))
	assert.True(t, strings.Contains(str, "login : sa"))
	assert.True(t, strings.Contains(str, "password : sa"))
	//assert.Equal(t, "connectTimeout : 1000, useSsl : true, login : sa, password : sa", str)

	anotherOptions := map[int]interface{}{
		1:  "value 1",
		2:  "value 2",
		-5: "value -5",
	}

	str = MapToString(&anotherOptions, KeyValueWithArrowSepFormat, ", ")
	assert.True(t, strings.Contains(str, "1 => value 1"))
	assert.True(t, strings.Contains(str, "2 => value 2"))
	assert.True(t, strings.Contains(str, "-5 => value -5"))
	//assert.Equal(t, "1 => value 1, 2 => value 2, -5 => value -5", str)
}
