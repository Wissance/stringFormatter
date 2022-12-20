package stringFormatter

import (
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, "connectTimeout : 1000, useSsl : true, login : sa, password : sa", str)
}
