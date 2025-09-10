package stringFormatter_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"testing"
)

func TestDefineFormattingStyle(t *testing.T) {
	for name, test := range map[string]struct {
		text     string
		expected string
		newStyle stringFormatter.FormattingStyle
	}{
		"snake-to-kebab-simple": {
			text:     "my_super_func",
			expected: "my-super-func",
			newStyle: stringFormatter.Kebab,
		},
		"kebab-to-snake-simple": {
			text:     "my-super-func",
			expected: "my_super_func",
			newStyle: stringFormatter.Snake,
		},
		"lower-case-camel-to-snake-simple": {
			text:     "mySuperFunc",
			expected: "my_super_func",
			newStyle: stringFormatter.Snake,
		},
		"snake-to-camel-simple": {
			text:     "my_super_func",
			expected: "mySuperFunc",
			newStyle: stringFormatter.Camel,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := stringFormatter.SetStyle(&test.text, test.newStyle)
			assert.Equal(t, test.expected, actual)
		})
	}
}
