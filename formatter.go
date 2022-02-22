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
/* Function that format text using more complex templates like "Hello {username} here is our application {appname}
 *
 */
func FormatComplex(template string, args map[string]interface{}) string {
	if args == nil {
		return template
	}
    errStr := ""
	formattedStr := template
	for key, val := range args {
		arg := "{" + key + "}"
		strVal, err := getItemAsStr(val)
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

// todo: umv: impl format passing as param
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
	    	case float32:
		    	strVal = strconv.FormatFloat(float64(item.(float32)), 'f', -1, 32)
		    	break
	    	case float64:
	    		strVal = strconv.FormatFloat(item.(float64), 'f', -1, 64)
	    		break
	     	default:
	     		strVal = fmt.Sprintf("%v", item)
	     		break
	}
	return strVal, err
}
