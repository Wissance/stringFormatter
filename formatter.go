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
func Format(template string, args ...any) string {
	if args == nil {
		return template
	}

	start := strings.Index(template, "{")
	if start < 0 {
		return template
	}

	templateLen := len(template)
	formattedStr := &strings.Builder{}
	formattedStr.Grow(templateLen + 22*len(args))
	j := -1 //nolint:ineffassign

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
			}
			// find end of placeholder
			j = i + 2
			for {
				if j >= templateLen {
					break
				}

				if template[j] == '{' {
					// multiple nested curly brackets ...
					formattedStr.WriteString(template[i:j])
					i = j
				}

				if template[j] == '}' {
					break
				}

				j++
			}
			// double curly brackets processed here, convert {{N}} -> {N}
			// so we catch here {{N}
			if j+1 < templateLen && template[j+1] == '}' && template[i-1] == '{' {
				formattedStr.WriteString(template[i+1 : j+1])
				i = j + 1
			} else {
				argNumberStr := template[i+1 : j]
				var argNumber int
				var err error
				if len(argNumberStr) == 1 {
					// this makes work a little faster than AtoI
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
			j = i //nolint:ineffassign
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
func FormatComplex(template string, args map[string]any) string {
	if args == nil {
		return template
	}

	start := strings.Index(template, "{")
	if start < 0 {
		return template
	}

	templateLen := len(template)
	formattedStr := &strings.Builder{}
	formattedStr.Grow(templateLen + 22*len(args))
	j := -1 //nolint:ineffassign
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
			}

			// find end of placeholder
			j = i + 2
			for {
				if j >= templateLen {
					break
				}
				if template[j] == '{' {
					// multiple nested curly brackets ...
					formattedStr.WriteString(template[i:j])
					i = j
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
					strVal := getItemAsStr(&arg)
					formattedStr.WriteString(strVal)
				} else {
					formattedStr.WriteByte('{')
					formattedStr.WriteString(argNumberStr)
					formattedStr.WriteByte('}')
				}
				i = j
			}
		} else {
			j = i //nolint:ineffassign
			formattedStr.WriteByte(template[i])
		}
	}

	return formattedStr.String()
}

// todo: umv: impl format passing as param
func getItemAsStr(item *any) string {
	switch v := (*item).(type) {
	case string:
		return v
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}
