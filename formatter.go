package stringFormatter

import (
	"fmt"
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

	templateLen := len(template)
	var formattedStr = &strings.Builder{}
	formattedStr.Grow(templateLen + 22*len(args))
	j := -1
	for i := 0; i < templateLen; i++ {

		if template[i] == '{' {
			// possibly it is a template placeholder
			if i == templateLen-1 {
				break
			}
			if template[i+1] == '{' { // todo: umv: this not considering {{0}}
				formattedStr.WriteByte('{')
				continue
			} else {
				// find end of placeholder
				j = i + 2
				for {
					if j >= templateLen {
						break
					}
					if template[j] == '}' {
						break
					}
					j++
				}

				argNumberStr := template[i+1 : j]
				var argNumber int
				var err error
				if len(argNumberStr) == 1 {
					// this makes work a little faster then AtoI
					argNumber = int(argNumberStr[0] - '0')
				} else {
					argNumber, err = strconv.Atoi(argNumberStr)
				}
				//argNumber, err := strconv.Atoi(argNumberStr)
				if err == nil && len(args) > argNumber {
					// get number from placeholder
					strVal := getItemAsStr(&args[argNumber])
					formattedStr.WriteString(strVal)
				} else {
					formattedStr.WriteByte('{')
					formattedStr.WriteString(argNumberStr)
					formattedStr.WriteByte('}')
				}
				i = j
			}

		} else {
			j = i
			formattedStr.WriteByte(template[i])
		}
	}

	return formattedStr.String()
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

	formattedStr := template
	for key, val := range args {
		arg := "{" + key + "}"
		strVal := getItemAsStr(&val)
		// todo: ignore double curly brackets
		formattedStr = strings.Replace(formattedStr, arg, strVal, -1)
	}

	return formattedStr
}

// todo: umv: impl format passing as param
func getItemAsStr(item *interface{}) string {
	var strVal string
	//var err error
	switch (*item).(type) {
	case string:
		strVal = (*item).(string)
		break
	case int8:
		strVal = strconv.FormatInt(int64((*item).(int8)), 10)
		break
	case int16:
		strVal = strconv.FormatInt(int64((*item).(int16)), 10)
		break
	case int32:
		strVal = strconv.FormatInt(int64((*item).(int32)), 10)
		break
	case int64:
		strVal = strconv.FormatInt((*item).(int64), 10)
		break
	case int:
		strVal = strconv.FormatInt(int64((*item).(int)), 10)
		break
	case uint8:
		strVal = strconv.FormatUint(uint64((*item).(uint8)), 10)
		break
	case uint16:
		strVal = strconv.FormatUint(uint64((*item).(uint16)), 10)
		break
	case uint32:
		strVal = strconv.FormatUint(uint64((*item).(uint32)), 10)
		break
	case uint64:
		strVal = strconv.FormatUint((*item).(uint64), 10)
		break
	case uint:
		strVal = strconv.FormatUint(uint64((*item).(uint)), 10)
		break
	case bool:
		strVal = strconv.FormatBool((*item).(bool))
		break
	case float32:
		strVal = strconv.FormatFloat(float64((*item).(float32)), 'f', -1, 32)
		break
	case float64:
		strVal = strconv.FormatFloat((*item).(float64), 'f', -1, 64)
		break
	default:
		strVal = fmt.Sprintf("%v", *item)
		break
	}
	return strVal
}
