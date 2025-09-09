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
	} {
		t.Run(name, func(t *testing.T) {
			actual := stringFormatter.SetStyle(&test.text, test.newStyle)
			assert.Equal(t, test.expected, actual)
		})
	}
}
