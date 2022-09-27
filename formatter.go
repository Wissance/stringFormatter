package stringFormatter

import (
	"fmt"
	"github.com/golang/glog"
	"strconv"
	"strings"
)

// Format
/* Func that makes string formatting from template
 * It differs from above function only by generic interface that allow to use only primitive data types:
 * - integers (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uin64)
 * - floats (float32, float64)
 * - boolean
 * - string
 * - complex
 * - objects
 * This function defines format automatically
 * Parameters
 *    - template - string that contains template
 *    - args - values that are using for formatting with template
 * Returns formatted string
 */
func Format(template string, args ...interface{}) string {
	if args == nil {
		return template
	}
	errStr := ""
	formattedStr := template
	for index, val := range args {
		arg := "{" + strconv.Itoa(index) + "}"
		strVal, err := getItemAsStr(&val)
		if err != nil {
			errStr += err.Error()
			errStr += "\n"
		}
		formattedStr = strings.Replace(formattedStr, arg, strVal, -1)
	}

	if len(errStr) > 0 {
		glog.Warning(errStr)
	}
	return formattedStr
}

// FormatComplex
/* Function that format text using more complex templates contains string literals i.e "Hello {username} here is our application {appname}
 * Parameters
 *    - template - string that contains template
 *    - args - values (dictionary: string key - any value) that are using for formatting with template
 * Returns formatted string
 */
func FormatComplex(template string, args map[string]interface{}) string {
	if args == nil {
		return template
	}
	errStr := ""
	formattedStr := template
	for key, val := range args {
		arg := "{" + key + "}"
		strVal, err := getItemAsStr(&val)
		if err != nil {
			errStr += err.Error()
			errStr += "\n"
		}
		// formattedStr = strVal + arg
		formattedStr = strings.Replace(formattedStr, arg, strVal, -1)
	}
	if len(errStr) > 0 {
		glog.Warning(errStr)
	}

	return formattedStr
}

// todo: umv: impl format passing as param
func getItemAsStr(item *interface{}) (string, error) {
	var strVal string
	var err error
	value := *item

	switch value.(type) {
	case int8:
		strVal = NumericToStr[int8](value.(int8), 10)
		break
	case int16:
		strVal = NumericToStr[int16](value.(int16), 10)
		break
	case int32:
		strVal = NumericToStr[int32](value.(int32), 10)
		break
	case int64:
		strVal = NumericToStr[int64](value.(int64), 10)
		break
	case int:
		strVal = NumericToStr[int](value.(int), 10)
		break
	case uint8:
		strVal = NumericToStr[uint8](value.(uint8), 10)
		break
	case uint16:
		strVal = NumericToStr[uint16](value.(uint16), 10)
		break
	case uint32:
		strVal = NumericToStr[uint32](value.(uint32), 10)
		break
	case uint64:
		strVal = NumericToStr[uint64](value.(uint64), 10)
		break
	case uint:
		strVal = NumericToStr[uint](value.(uint), 10)
		break
	case string:
		strVal = value.(string)
		break
	case bool:
		strVal = strconv.FormatBool(value.(bool))
		break
	case float32:
		strVal = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32)
		break
	case float64:
		strVal = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		break
	default:
		strVal = fmt.Sprintf("%v", *item)
		break
	}
	return strVal, err
}
