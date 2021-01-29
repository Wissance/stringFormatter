package stringFormatter

import (
	"fmt"
	"strconv"
	"strings"
)

/* Func that makes string formatting from template
 * It differs from above function only by generic interface that allow to use only primitive data types:
 * - integers (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uin64)
 * - floats
 * - boolean
 * - string
 * - complex
 */
func Format(template string, args ...interface{}) string {
	if args == nil {
		return template
	}
	errStr := ""
	formattedStr := template
	for index, val := range args {
		arg := "{" + strconv.Itoa(index) + "}"
		strVal, err := getItemAsStr(val)
		fmt.Println(strVal)
		if err != nil {
			errStr += err.Error()
			errStr += "\n"
		}

		formattedStr = strings.Replace(formattedStr, arg, strVal, -1)
	}
	/*var err error = nil
	if len(errStr) > 0 {
		err = errors.New(errStr)
	}*/
	return formattedStr
}

/* Func that format more complex templates like "Hello {username} here is our application {appname}
 *
 */
func FormatComplex(template string, args map[string]string)  (string, error) {
	if args == nil {
		return template, nil
	}

	formattedStr := template
	for key, val := range args {
		arg := "{" + key + "}"
		formattedStr = strings.Replace(formattedStr, arg, val, -1)
	}
    return formattedStr, nil
}

// todo: umv: impl format passing as param
// todo: add complex support
func getItemAsStr(item interface{}) (string, error) {
    var strVal string
    var err error
    switch item.(type) {
	    case int8:
			strVal = strconv.FormatInt(int64(item.(int8)), 10)
			break
	    case int16:
		    strVal = strconv.FormatInt(int64(item.(int16)), 10)
		    break
	    case int32:
		    strVal = strconv.FormatInt(int64(item.(int32)), 10)
		    break
	    case int64:
		    strVal = strconv.FormatInt(item.(int64), 10)
		    break
	    case int:
		    strVal = strconv.FormatInt(int64(item.(int)), 10)
		    break
		case uint8:
			strVal = strconv.FormatUint(uint64(item.(uint8)), 10)
			break
		case uint16:
			strVal = strconv.FormatUint(uint64(item.(uint16)), 10)
			break
		case uint32:
			strVal = strconv.FormatUint(uint64(item.(uint32)), 10)
			break
		case uint64:
			strVal = strconv.FormatUint(item.(uint64), 10)
			break
		case uint:
			strVal = strconv.FormatUint(uint64(item.(uint)), 10)
			break
	    case string:
		    strVal = item.(string)
		    break
	    case bool:
	    	strVal = strconv.FormatBool(item.(bool))
	    	break
	    case float32, float64:
	    	strVal = strconv.FormatFloat(item.(float64), 'f', -1, 64)
	    	break
	}
	return strVal, err
}