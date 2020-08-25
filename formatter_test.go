package stringFormatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrFormat(t *testing.T) {
    strFormatResult, err := Format("Hello i am {0}, my age is {1} and i am waiting for {2}, because i am {0}!",
    	                      "Michael Ushakov (Evillord666)", "34", "\"Great Success\"")
    assert.Nil(t, err)
    assert.Equal(t, "Hello i am Michael Ushakov (Evillord666), my age is 34 and i am waiting for \"Great Success\", because i am Michael Ushakov (Evillord666)!", strFormatResult)

    strFormatResult, err = Format("We are wondering if these values would be replaced : {5}, {4}, {0}", "one", "two", "three")
	assert.Nil(t, err)
	assert.Equal(t, "We are wondering if these values would be replaced : {5}, {4}, one", strFormatResult)

	strFormatResult, err = Format("No args ... : {0}, {1}, {2}")
	assert.Nil(t, err)
	assert.Equal(t, "No args ... : {0}, {1}, {2}", strFormatResult)
}