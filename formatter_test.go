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
			template: `{"StartAt": "S0", "States": {"S0": {"Type": "Map" {0}, ` +
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
		"": {
			template: "This is the text with an only number formatting: decimal - {0} / {0 : D6}, scientific - {1} / {1 : e2}",
			args:     []any{123, 191.0784},
			expected: "This is the text with an only number formatting: decimal - 123 / 000123, scientific - 191.0784 / 1.91e+02",
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
		"format json": {
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
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, stringFormatter.FormatComplex(test.template, test.args))
		})
	}
}
