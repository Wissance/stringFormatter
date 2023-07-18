package stringFormatter_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wissance/stringFormatter"
)

const _separator = ", "

func TestMapToString(t *testing.T) {
	for name, test := range map[string]struct {
		str           string
		expectedParts []string
	}{
		"semicolon sep": {
			str: stringFormatter.MapToString(
				map[string]any{
					"connectTimeout": 1000,
					"useSsl":         true,
					"login":          "sa",
					"password":       "sa",
				},
				"{key} : {value}",
				_separator,
			),
			expectedParts: []string{
				"connectTimeout : 1000",
				"useSsl : true",
				"login : sa",
				"password : sa",
			},
		},
		"arrow sep": {
			str: stringFormatter.MapToString(
				map[int]any{
					1:  "value 1",
					2:  "value 2",
					-5: "value -5",
				},
				"{key} => {value}",
				_separator,
			),
			expectedParts: []string{
				"1 => value 1",
				"2 => value 2",
				"-5 => value -5",
			},
		},
		"only value": {
			str: stringFormatter.MapToString(
				map[uint64]any{
					1: "value 1",
					2: "value 2",
					5: "value 5",
				},
				"{value}",
				_separator,
			),
			expectedParts: []string{
				"value 1",
				"value 2",
				"value 5",
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actualParts := strings.Split(test.str, _separator)
			assert.ElementsMatch(t, test.expectedParts, actualParts)
		})
	}
}
