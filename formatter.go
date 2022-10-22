package stringFormatter

import (
	"fmt"
	"strconv"
	"strings"
)

type state byte

const (
	stateText state = iota
	stateInCurly
	stateEndCurly
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
	formattedStr := &strings.Builder{}
	formattedStr.Grow(templateLen + 22*len(args))

	last := 0
	state := stateText

	for i, c := range template {
		switch state {
		case stateText:
			switch c {
			case '{':
				formattedStr.WriteString(template[last:i])
				last = i
				state = stateInCurly
			case '}':
				formattedStr.WriteString(template[last:i])
				last = i
				state = stateEndCurly
			default:
			}
		case stateEndCurly:
			if c != '}' {
				return "error, but there is no error return parameter"
			}
			formattedStr.WriteRune('}')
			last = i + 1
			state = stateText
		case stateInCurly:
			// Safety check but old code doesn't have it so I won't waste cycles here
			// if c < '0' || c > '9' {
			// 	return "error, but there is no error return parameter"
			// }
			switch c {
			case '}':
				argNumberStr := template[last+1 : i]

				var argNumber int
				var err error
				if len(argNumberStr) == 1 {
					// this makes work a little faster then AtoI
					argNumber = int(argNumberStr[0] - '0')
				} else {
					argNumber, err = strconv.Atoi(argNumberStr)
				}

				if err == nil && len(args) > argNumber {
					// get number from placeholder
					strVal := getItemAsStr(args[argNumber])
					formattedStr.WriteString(strVal)
				} else {
					formattedStr.WriteString(template[last : i+1])
				}

				last = i + 1
				state = stateText
			case '{':
				if last+1 != i {
					return "error, but there is no error return parameter"
				}
				formattedStr.WriteByte('{')
				last = i + 1
				state = stateText
			default:
			}
		}
	}

	switch state {
	case stateText:
		formattedStr.WriteString(template[last:])
		return formattedStr.String()
	case stateInCurly, stateEndCurly:
		return "error, but there is no error return parameter"
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

	templateLen := len(template)
	var formattedStr = &strings.Builder{}
	formattedStr.Grow(templateLen + 22*len(args))
	j := -1
	start := strings.Index(template, "{")
	if start < 0 {
		return template
	}

	formattedStr.WriteString(template[:start])
	for i := start; i < templateLen; i++ {

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
				// double curly brackets processed here, convert {{N}} -> {N}
				// so we catch here {{N}
				if j+1 < templateLen && template[j+1] == '}' {

					formattedStr.WriteString(template[i+1 : j+1])
					i = j + 1
				} else {
					argNumberStr := template[i+1 : j]
					arg, ok := args[argNumberStr]
					if ok {
						// get number from placeholder
						strVal := getItemAsStr(arg)
						formattedStr.WriteString(strVal)
					} else {
						formattedStr.WriteByte('{')
						formattedStr.WriteString(argNumberStr)
						formattedStr.WriteByte('}')
					}
					i = j
				}
			}

		} else {
			j = i
			formattedStr.WriteByte(template[i])
		}
	}

	return formattedStr.String()
}

// todo: umv: impl format passing as param
func getItemAsStr(item interface{}) string {
	var strVal string
	//var err error
	switch i := item.(type) {
	case string:
		strVal = item.(string)
	case int8:
		strVal = strconv.FormatInt(int64(i), 10)
	case int16:
		strVal = strconv.FormatInt(int64(i), 10)
	case int32:
		strVal = strconv.FormatInt(int64(i), 10)
	case int64:
		strVal = strconv.FormatInt(i, 10)
	case int:
		strVal = strconv.FormatInt(int64(i), 10)
	case uint8:
		strVal = strconv.FormatUint(uint64(i), 10)
	case uint16:
		strVal = strconv.FormatUint(uint64(i), 10)
	case uint32:
		strVal = strconv.FormatUint(uint64(i), 10)
	case uint64:
		strVal = strconv.FormatUint(i, 10)
	case uint:
		strVal = strconv.FormatUint(uint64(i), 10)
	case bool:
		strVal = strconv.FormatBool(i)
	case float32:
		strVal = strconv.FormatFloat(float64(i), 'f', -1, 32)
	case float64:
		strVal = strconv.FormatFloat(i, 'f', -1, 64)
	default:
		strVal = fmt.Sprintf("%v", item)
	}
	return strVal
}
