package stringFormatter

import (
	"errors"
	"strconv"
	"strings"
)

/* Func that make string formatting from template
 * Template is a following "Hello {0}, It is time to go to the {1}, {0}!"
 * if you are passing to upper mentioned str template values John, library you get
 * following result: Hello John, It is time to go to the library, John!
 */
func Format(template string, args ...string) (string, error){

	// we return here original string
	if args == nil {
		return template, nil
	}

	formattedStr := template

	for index, arg := range args{
		argNr := "{" + strconv.Itoa(index) + "}"
		formattedStr = strings.Replace(formattedStr, argNr, arg, -1)
	}

	return formattedStr, nil
}

/* Func that makes string formatting from template
 * It differs from above function only by generic interface that allow to use only primitive data types:
 * - integers
 * - floats
 * - boolean
 * - string
 * - complex
 */
func FormatGeneric(template string, args ...interface{}) (string, error) {
	if args == nil {
		return template, nil
	}
	errStr := ""
	formattedStr := template
	for index, val := range args {
		arg := "{" + strconv.Itoa(index) + "}"
		strVal, err := getItemAsStr(val)
		errStr += err.Error()
		errStr += "\n"
		formattedStr = strings.Replace(formattedStr, arg, strVal, -1)
	}
	return formattedStr, errors.New(errStr)
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
func getItemAsStr(item interface{}) (string, error) {
    var strVal string
    var err error
    switch item.(type) {
	    case int:
	    case int8:
	    case int16:
	    case int32:
	    case int64:
	    	strVal = strconv.FormatInt(item.(int64), 10)
	    	break
	    case uint:
	    case uint8:
	    case uint16:
	    case uint32:
	    case uint64:
		    strVal = strconv.FormatUint(item.(uint64), 10)
		    break
	    case string:
		    strVal = item.(string)
		    break
	    case bool:
	    	strVal = strconv.FormatBool(item.(bool))
	    	break
	    case float32:
	    case float64:
	    	strVal = strconv.FormatFloat(item.(float64), 'f', -1, 64)
	    	break
	}
	return strVal, err
}