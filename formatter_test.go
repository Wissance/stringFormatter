package stringFormatter_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wissance/stringFormatter"
)

const _address = "grpcs://127.0.0.1"

type meteoData struct {
	Int    int
	Str    string
	Double float64
	Err    error
}

func TestFormat(t *testing.T) {
	for name, test := range map[string]struct {
		template string
		args     []any
		expected string
	}{
		"all args in place": {
			template: "Hello i am {0}, my age is {1} and i am waiting for {2}, because i am {0}!",
			args:     []any{"Michael Ushakov (Evillord666)", "34", `"Great Success"`},
			expected: `Hello i am Michael Ushakov (Evillord666), my age is 34 and i am waiting for "Great Success", because i am Michael Ushakov (Evillord666)!`,
		},
		"too large index": {
			template: "We are wondering if these values would be replaced : {5}, {4}, {0}",
			args:     []any{"one", "two", "three"},
			expected: "We are wondering if these values would be replaced : {5}, {4}, one",
		},
		"no args": {
			template: "No args ... : {0}, {1}, {2}",
			args:     nil,
			expected: "No args ... : {0}, {1}, {2}",
		},
		"format json": {
			template: `
		    {
		         "Comment": "Call Lambda with GRPC",
		         "StartAt": "CallLambdaWithGrpc",
		         "States": {"CallLambdaWithGrpc": {"Type": "Task", "Resource": "{0}:get ad user", "End": true}}
		    }`,
			args: []any{_address},
			expected: `
		    {
		         "Comment": "Call Lambda with GRPC",
		         "StartAt": "CallLambdaWithGrpc",
		         "States": {"CallLambdaWithGrpc": {"Type": "Task", "Resource": "grpcs://127.0.0.1:get ad user", "End": true}}
		    }`,
		},
		"multiple nested curly brackets": {
			template: `{"StartAt": "S0", "States": {"S0": {"Type": "Map" , ` +
				`"Iterator": {"StartAt": "SI0", "States": {"SI0": {"Type": "Pass", "End": true}}}` +
				`, "End": true}}}`,
			args:     []any{""},
			expected: `{"StartAt": "S0", "States": {"S0": {"Type": "Map" , "Iterator": {"StartAt": "SI0", "States": {"SI0": {"Type": "Pass", "End": true}}}, "End": true}}}`,
		},
		"indexes out of args range": {
			template: "{3} - rings to the immortal elfs, {7} to dwarfs, {9} to greedy people and {1} to control everything",
			args:     []any{"0", "1", "2", "3"},
			expected: "3 - rings to the immortal elfs, {7} to dwarfs, {9} to greedy people and 1 to control everything",
		},
		"format integers": {
			template: `Here we are testing integers "int8": {0}, "int16": {1}, "int32": {2}, "int64": {3} and finally "int": {4}`,
			args:     []any{int8(8), int16(-16), int32(32), int64(-64), int(123)},
			expected: `Here we are testing integers "int8": 8, "int16": -16, "int32": 32, "int64": -64 and finally "int": 123`,
		},
		"format unsigneds": {
			template: `Here we are testing integers "uint8": {0}, "uint16": {1}, "uint32": {2}, "uint64": {3} and finally "uint": {4}`,
			args:     []any{uint8(8), uint16(16), uint32(32), uint64(64), uint(128)},
			expected: `Here we are testing integers "uint8": 8, "uint16": 16, "uint32": 32, "uint64": 64 and finally "uint": 128`,
		},
		"format floats": {
			template: `Here we are testing floats "float32": {0}, "float64":{1}`,
			args:     []any{float32(1.24), float64(1.56)},
			expected: `Here we are testing floats "float32": 1.24, "float64":1.56`,
		},
		"format bools": {
			template: `Here we are testing "bool" args: {0}, {1}`,
			args:     []any{false, true},
			expected: `Here we are testing "bool" args: false, true`,
		},
		"format complex": {
			template: `Here we are testing "complex64" {0} and "complex128": {1}`,
			args:     []any{complex64(complex(1.0, 6.0)), complex(2.3, 3.2)},
			expected: `Here we are testing "complex64" (1+6i) and "complex128": (2.3+3.2i)`,
		},
		"doubly curly brackets": {
			template: "Hello i am {{0}}, my age is {1} and i am waiting for {2}, because i am {0}!",
			args:     []any{"Michael Ushakov (Evillord666)", "34", `"Great Success"`},
			expected: `Hello i am {0}, my age is 34 and i am waiting for "Great Success", because i am Michael Ushakov (Evillord666)!`,
		},
		"doubly curly brackets at the end": {
			template: "At the end {{0}}",
			args:     []any{"s"},
			expected: "At the end {0}",
		},
		"quadro curly brackets in the middle": {
			template: "Not at the end {{{{0}}}}, in the middle",
			args:     []any{"s"},
			expected: "Not at the end {{0}}, in the middle",
		},
		"struct arg": {
			template: "Example is: {0}",
			args: []any{
				meteoData{
					Int:    123,
					Str:    "This is a test str, nothing more special",
					Double: -1.098743,
					Err:    errors.New("main question error, is 42"),
				},
			},
			expected: "Example is: {123 This is a test str, nothing more special -1.098743 main question error, is 42}",
		},
		"open bracket at the end of line of go line": {
			template: "type serviceHealth struct {",
			args:     []any{},
			expected: "type serviceHealth struct {",
		},
		"open bracket at the end of line of go line with {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) {",
			args:     []any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) {",
		},

		"close bracket at the end of line of go line with {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) }",
			args:     []any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) }",
		},

		"no bracket at the end of line with {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) ",
			args:     []any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) ",
		},
		"open bracket at the end of line of go line with multiple {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}, additionalData interface{}) {",
			args:     []any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}, additionalData interface{}) {",
		},
		"commentaries after bracket": {
			template: "switch app.appConfig.ServerCfg.Schema { //nolint:exhaustive",
			args:     []any{},
			expected: "switch app.appConfig.ServerCfg.Schema { //nolint:exhaustive",
		},
		"bracket in the middle": {
			template: "in the middle - { at the end - nothing",
			args:     []any{},
			expected: "in the middle - { at the end - nothing",
		},
		"code line with interface": {
			template: "[]any{singleValue}",
			args:     []any{},
			expected: "[]any{singleValue}",
		},
		"code line with interface with val": {
			template: "[]any{{{0}}}",
			args:     []any{"\"USSR!\""},
			expected: "[]any{\"USSR!\"}",
		},
		"2-symb str": {
			template: "a}",
			args:     []any{},
			expected: "a}",
		},
		"one symb segment": {
			template: "{x}",
			args:     []any{},
			expected: "{x}",
		},
		"one symb template": {
			template: "{",
			args:     []any{},
			expected: "{",
		},
		"one symb template2": {
			template: "}",
			args:     []any{},
			expected: "}",
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, stringFormatter.Format(test.template, test.args...))
		})
	}
}

func TestFormatWithArgFormatting(t *testing.T) {
	for name, test := range map[string]struct {
		template string
		args     []any
		expected string
	}{
		"numeric_test_1": {
			template: "This is the text with an only number formatting: decimal - {0} / {0 : D6}, scientific - {1} / {1 : e2}",
			args:     []any{123, 191.0784},
			expected: "This is the text with an only number formatting: decimal - 123 / 000123, scientific - 191.0784 / 1.91e+02",
		},
		"numeric_test_2": {
			template: "This is the text with an only number formatting: binary - {0:B} / {0 : B8}, hexadecimal - {1:X} / {1 : X4}",
			args:     []any{15, 250},
			expected: "This is the text with an only number formatting: binary - 1111 / 00001111, hexadecimal - fa / 00fa",
		},
		"numeric_test_3": {
			template: "This is the text with an only number formatting: decimal - {0:F} / {0 : F4} / {0:F8}",
			args:     []any{10.5467890},
			expected: "This is the text with an only number formatting: decimal - 10.546789 / 10.5468 / 10.54678900",
		},
		"numeric_test_4": {
			template: "This is the text with percentage format - {0:P100} / {0 : P100.5}, and non normal percentage {1:P100}",
			args:     []any{12, "ass"},
			expected: "This is the text with percentage format - 12.00 / 11.94, and non normal percentage 0.00",
		},
		"list_with_default_sep": {
			template: "This is a list(slice) test: {0:L}",
			args:     []any{[]any{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1,s2,s3",
		},
		"list_with_dash_sep": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]any{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1-s2-s3",
		},
		"list_with_space_sep": {
			template: "This is a list(slice) test: {0:L }",
			args:     []any{[]any{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1 s2 s3",
		},
	} {
		// Run test here
		t.Run(name, func(t *testing.T) {
			// assert.NotNil(t, test)
			assert.Equal(t, test.expected, stringFormatter.Format(test.template, test.args...))
		})
	}
}

func TestFormatWithArgFormattingForTypedSlice(t *testing.T) {
	for name, test := range map[string]struct {
		template string
		args     []any
		expected string
	}{
		"list_with_int_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]int{101, 202, 303}},
			expected: "This is a list(slice) test: 101-202-303",
		},
		"list_with_uint_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]uint{102, 204, 308}},
			expected: "This is a list(slice) test: 102-204-308",
		},
		"list_with_int32_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]int32{100, 200, 300}},
			expected: "This is a list(slice) test: 100-200-300",
		},
		"list_with_int64_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]int64{1001, 2002, 3003}},
			expected: "This is a list(slice) test: 1001-2002-3003",
		},
		"list_with_float64_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]float64{1.01, 2.02, 3.03}},
			expected: "This is a list(slice) test: 1.01-2.02-3.03",
		},
		"list_with_float32_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]float32{5.01, 6.02, 7.03}},
			expected: "This is a list(slice) test: 5.01-6.02-7.03",
		},
		"list_with_bool_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]bool{true, true, false}},
			expected: "This is a list(slice) test: true-true-false",
		},
		"list_with_string_slice": {
			template: "This is a list(slice) test: {0:L-}",
			args:     []any{[]string{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1-s2-s3",
		},
	} {
		// Run test here
		t.Run(name, func(t *testing.T) {
			// assert.NotNil(t, test)
			assert.Equal(t, test.expected, stringFormatter.Format(test.template, test.args...))
		})
	}
}

// TestStrFormatWithComplicatedText - this test represents issue with complicated text
func TestFormatComplex(t *testing.T) {
	for name, test := range map[string]struct {
		template string
		args     map[string]any
		expected string
	}{
		"numeric_test_1": {
			template: `
			{
				"Comment": "Call Lambda with GRPC",
				"StartAt": "CallLambdaWithGrpc",
				"States": {"CallLambdaWithGrpc": {"Type": "Task", "Resource": "{address}:get ad user", "End": true}}
			}`,
			args: map[string]any{"address": _address},
			expected: `
			{
				"Comment": "Call Lambda with GRPC",
				"StartAt": "CallLambdaWithGrpc",
				"States": {"CallLambdaWithGrpc": {"Type": "Task", "Resource": "grpcs://127.0.0.1:get ad user", "End": true}}
			}`,
		},
		"key not found": {
			template: "Hello: {username}, you earn {amount} $",
			args:     map[string]any{"amount": 1000},
			expected: "Hello: {username}, you earn 1000 $",
		},
		"dialog": {
			template: "Hello {user} what are you doing here {app} ?",
			args:     map[string]any{"user": "vpupkin", "app": "mn_console"},
			expected: "Hello vpupkin what are you doing here mn_console ?",
		},
		"info message": {
			template: "Current app settings are: ipAddr: {ipaddr}, port: {port}, use ssl: {ssl}.",
			args:     map[string]any{"ipaddr": "127.0.0.1", "port": 5432, "ssl": false},
			expected: "Current app settings are: ipAddr: 127.0.0.1, port: 5432, use ssl: false.",
		},
		"one json line with open bracket at the end": {
			template: "    \"server\": {",
			args:     map[string]any{},
			expected: "    \"server\": {",
		},
		"open bracket at the end of line of go line with {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) {",
			args:     map[string]any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) {",
		},

		"open bracket at the end of line of go line with multiple {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}, additionalData interface{}) {",
			args:     map[string]any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}, additionalData interface{}) {",
		},

		"close bracket at the end of line of go line with {} inside": {
			template: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) }",
			args:     map[string]any{},
			expected: "func afterHandle(respWriter *http.ResponseWriter, statusCode int, data interface{}) }",
		},
		"commentaries after bracket": {
			template: "switch app.appConfig.ServerCfg.Schema { //nolint:exhaustive",
			args:     map[string]any{},
			expected: "switch app.appConfig.ServerCfg.Schema { //nolint:exhaustive",
		},
		"code line with interface": {
			template: "[]any{singleValue}",
			args:     map[string]any{},
			expected: "[]any{singleValue}",
		},
		"code line with interface with val": {
			template: "[]any{{val}}",
			args:     map[string]any{"val": "\"USSR!\""},
			expected: "[]any{\"USSR!\"}",
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, stringFormatter.FormatComplex(test.template, test.args))
		})
	}
}

func TestFormatComplexWithArgFormatting(t *testing.T) {
	for name, test := range map[string]struct {
		template string
		args     map[string]any
		expected string
	}{
		"numeric_test_1": {
			template: "This is the text with an only number formatting: scientific - {mass} / {mass : e2}",
			args:     map[string]any{"mass": 191.0784},
			expected: "This is the text with an only number formatting: scientific - 191.0784 / 1.91e+02",
		},
		"numeric_test_2": {
			template: "This is the text with an only number formatting: binary - {bin:B} / {bin : B8}, hexadecimal - {hex:X} / {hex : X4}",
			args:     map[string]any{"bin": 15, "hex": 250},
			expected: "This is the text with an only number formatting: binary - 1111 / 00001111, hexadecimal - fa / 00fa",
		},
		"numeric_test_3": {
			template: "This is the text with an only number formatting: decimal - {float:F} / {float : F4} / {float:F8}",
			args:     map[string]any{"float": 10.5467890},
			expected: "This is the text with an only number formatting: decimal - 10.546789 / 10.5468 / 10.54678900",
		},
		"list_with_default_sep": {
			template: "This is a list(slice) test: {list:L}",
			args:     map[string]any{"list": []any{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1,s2,s3",
		},
		"list_with_dash_sep": {
			template: "This is a list(slice) test: {list:L-}",
			args:     map[string]any{"list": []any{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1-s2-s3",
		},
		"list_with_space_sep": {
			template: "This is a list(slice) test: {list:L }",
			args:     map[string]any{"list": []any{"s1", "s2", "s3"}},
			expected: "This is a list(slice) test: s1 s2 s3",
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, stringFormatter.FormatComplex(test.template, test.args))
		})
	}
}
