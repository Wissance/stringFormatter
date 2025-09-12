package stringFormatter_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"testing"
)

func TestSetFormattingStyleWithoutCaseModification(t *testing.T) {
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
			expected: "my_Super_Func",
			newStyle: stringFormatter.Snake,
		},
		"snake-to-camel-simple": {
			text:     "my_super_func",
			expected: "mySuperFunc",
			newStyle: stringFormatter.Camel,
		},
		"camel-to-snake-with-underscore-the-end": {
			text:     "myVal_",
			expected: "my_Val_",
			newStyle: stringFormatter.Snake,
		},
		"mixed-to-camel-simple": {
			text:     "TestGetManyMethod_WithDefaultParams",
			expected: "TestGetManyMethodWithDefaultParams",
			newStyle: stringFormatter.Camel,
		},
		"no-changes-simple": {
			text:     "_my_variable",
			expected: "_my_variable",
			newStyle: stringFormatter.Snake,
		},
		"snake_to_camel_with_underscore_at_start": {
			text:     "_my_variable",
			expected: "MyVariable",
			newStyle: stringFormatter.Camel,
		},
		"camel-with_abbreviation-to-snake": {
			text:     "convertToJSON",
			expected: "convert_To_JSON",
			newStyle: stringFormatter.Snake,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := stringFormatter.SetStyle(&test.text, test.newStyle, stringFormatter.NoChanges,
				stringFormatter.NoChanges)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSetFormattingStyleWithCaseModification(t *testing.T) {
	for name, test := range map[string]struct {
		text            string
		expected        string
		newStyle        stringFormatter.FormattingStyle
		firstSymbolCase stringFormatter.CaseSetting
		textCase        stringFormatter.CaseSetting
	}{
		"snake-to-kebab-upper-case": {
			text:            "my_super_func",
			expected:        "MY-SUPER-FUNC",
			newStyle:        stringFormatter.Kebab,
			firstSymbolCase: stringFormatter.ToUpper,
			textCase:        stringFormatter.ToUpper,
		},
		"snake-to-camel-starting-from-lower-case": {
			text:            "my_super_func",
			expected:        "mySuperFunc",
			newStyle:        stringFormatter.Camel,
			firstSymbolCase: stringFormatter.ToLower,
			textCase:        stringFormatter.NoChanges,
		},
		"camel-to-upper-case-snake": {
			text:            "mySuperFunc",
			expected:        "MY_SUPER_FUNC",
			newStyle:        stringFormatter.Snake,
			firstSymbolCase: stringFormatter.ToUpper,
			textCase:        stringFormatter.ToUpper,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := stringFormatter.SetStyle(&test.text, test.newStyle, test.firstSymbolCase,
				test.textCase)
			assert.Equal(t, test.expected, actual)
		})
	}
}
