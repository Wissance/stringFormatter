package stringFormatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrFormat(t *testing.T) {
    strFormatResult := Format("Hello i am {0}, my age is {1} and i am waiting for {2}, because i am {0}!",
    	                      "Michael Ushakov (Evillord666)", "34", "\"Great Success\"")
    assert.Equal(t, "Hello i am Michael Ushakov (Evillord666), my age is 34 and i am waiting for \"Great Success\", because i am Michael Ushakov (Evillord666)!", strFormatResult)

    strFormatResult = Format("We are wondering if these values would be replaced : {5}, {4}, {0}", "one", "two", "three")
	assert.Equal(t, "We are wondering if these values would be replaced : {5}, {4}, one", strFormatResult)

	strFormatResult = Format("No args ... : {0}, {1}, {2}")
	assert.Equal(t, "No args ... : {0}, {1}, {2}", strFormatResult)
}

func TestStrFormatGeneric(t *testing.T) {
	strFormat1 := "Here we are testing integers \"int8\": {0}, \"int16\": {1}, \"int32\": {2}, \"int64\": {3} and finally \"int\": {4}"
	var v1 int8 = 8
	var v2 int16 = -16
	var v3 int32 = 32
	var v4 int64 = -64
	var v5 int = 123

	strFormatResult := Format(strFormat1, v1, v2, v3, v4, v5)
	assert.Equal(t, "Here we are testing integers \"int8\": 8, \"int16\": -16, \"int32\": 32, \"int64\": -64 and finally \"int\": 123", strFormatResult)

	strFormat2 := "Here we are testing integers \"uint8\": {0}, \"uint16\": {1}, \"uint32\": {2}, \"uint64\": {3} and finally \"uint\": {4}"
	var v6 uint8 = 8
	var v7 uint16 = 16
	var v8 uint32 = 32
	var v9 uint64 = 64
	var v10 uint = 128

	strFormatResult = Format(strFormat2, v6, v7, v8, v9, v10)
	assert.Equal(t, "Here we are testing integers \"uint8\": 8, \"uint16\": 16, \"uint32\": 32, \"uint64\": 64 and finally \"uint\": 128", strFormatResult)

	strFormat3 := "Here we are testing floats \"float32\": {0}, \"float64\":{1}"
	var v11 float32 = 1.24
	var v12 float64 = 1.56
	strFormatResult = Format(strFormat3, v11, v12)
	assert.Equal(t, "Here we are testing floats \"float32\": 1.24, \"float64\":1.56", strFormatResult)

	strFormat4 := "Here we are testing \"bool\" args: {0}, {1}"
	var v13 bool = false
	var v14 bool = true
	strFormatResult = Format(strFormat4, v13, v14)
	assert.Equal(t, "Here we are testing \"bool\" args: false, true", strFormatResult)
}

func TestStrFormatComplex(t *testing.T) {
	strFormatResult,_ := FormatComplex("Hello {user} what are you doing here {app} ?", map[string]string{"user":"vpupkin", "app":"mn_console"})
	//assert.Nil(t, err)
	assert.Equal(t, "Hello vpupkin what are you doing here mn_console ?", strFormatResult)
}