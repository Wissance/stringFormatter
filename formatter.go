package stringFormatter

import (
	"fmt"
	"strconv"
	"strings"
)

const argumentFormatSeparator = ":"

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

	nestedBrackets := false
	formattedStr.WriteString(template[:start])
	for i := start; i < templateLen; i++ {
		if template[i] == '{' {
			// possibly it is a template placeholder
			if i == templateLen-1 {
				break
			}
			// considering in 2 phases - {{ }}
			if template[i+1] == '{' {
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
					nestedBrackets = true
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
				// is here we should support formatting ?
				var argNumber int
				var err error
				var argFormatOptions string
				if len(argNumberStr) == 1 {
					// this calculation makes work a little faster than AtoI
					argNumber = int(argNumberStr[0] - '0')
				} else {
					argNumber = -1
					// Here we are going to process argument either with additional formatting or not
					// i.e. 0 for arg without formatting && 0:format for an argument wit formatting
					// todo(UMV): we could format json or yaml here ...
					formatOptionIndex := strings.Index(argNumberStr, argumentFormatSeparator)
					// formatOptionIndex can't be == 0, because 0 is a position of arg number
					if formatOptionIndex > 0 {
						// trim formatting string
						argParts := strings.Split(argNumberStr, argumentFormatSeparator)
						argFormatOptions = strings.Trim(argParts[1], " ")
						argNumber, err = strconv.Atoi(strings.Trim(argParts[0], " "))
						if err == nil {
							argNumberStr = argParts[0]
						}
						// make formatting option str for further pass to an argument
					}
					//
					if argNumber < 0 {
						argNumber, err = strconv.Atoi(argNumberStr)
					}
				}
				//_, err2 := strconv.Atoi(argNumberStr)
				if (err == nil || (argFormatOptions != "" && !nestedBrackets)) &&
					len(args) > argNumber {
					// get number from placeholder
					strVal := getItemAsStr(&args[argNumber], &argFormatOptions)
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
	nestedBrackets := false
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
					nestedBrackets = true
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
				var argFormatOptions string
				argNumberStr := template[i+1 : j]
				arg, ok := args[argNumberStr]
				if !ok {
					argParts := strings.Split(argNumberStr, argumentFormatSeparator)
					if len(argParts) > 1 {
						argFormatOptions = strings.Trim(argParts[1], " ")
						argNumberStr = strings.Trim(argParts[0], " ")
					}
					arg, ok = args[argNumberStr]
				}
				if ok || (argFormatOptions != "" && !nestedBrackets) {
					// get number from placeholder
					strVal := getItemAsStr(&arg, &argFormatOptions)
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

func getItemAsStr(item *any, itemFormat *string) string {
	base := 10
	var floatFormat byte = 'f'
	precision := -1
	var preparedArgFormat string
	var argStr string
	postProcessingRequired := false
	intNumberFormat := false
	floatNumberFormat := false

	if itemFormat != nil && len(*itemFormat) > 0 {
		/* for numbers there are following formats:
		 * d(D) - decimal
		 * b(B) - binary
		 * f(F) - fixed point i.e {0:F}, 10.5467890 -> 10.546789 ; {0:F4}, 10.5467890 -> 10.5468
		 * e(E) - exponential - float point with scientific format {0:E2}, 191.0784 -> 1.91e+02
		 * x(X) - hexadecimal i.e. {0:X}, 250 -> fa ; {0:X4}, 250 -> 00fa
		 * p(P) - percent i.e. {0:P100}, 12 -> 12%
		 * Following formats are not supported yet:
		 *   1. c(C) currency it requires also country code
		 *   2. g(G),and others with locales
		 *   3. f(F) - fixed point, {0,F4}, 123.15 -> 123.1500
		 * OUR own addition:
		 * 1. O(o) - octahedral number format
		 */
		preparedArgFormat = strings.ToLower(*itemFormat)
		postProcessingRequired = len(preparedArgFormat) > 1

		switch rune(preparedArgFormat[0]) {
		case 'd':
			base = 10
			intNumberFormat = true
		case 'x':
			base = 16
			intNumberFormat = true
		case 'o':
			base = 8
			intNumberFormat = true
		case 'b':
			base = 2
			intNumberFormat = true
		case 'e', 'f':
			if rune(preparedArgFormat[0]) == 'e' {
				floatFormat = 'e'
			}
			// precision was passed, take [1:end], extract precision
			if postProcessingRequired {
				precisionStr := preparedArgFormat[1:]
				precisionVal, err := strconv.Atoi(precisionStr)
				if err == nil {
					precision = precisionVal
				}
			}
			postProcessingRequired = false
			floatNumberFormat = preparedArgFormat[0] == 'f'

		case 'p':
			// percentage processes here ...
			if postProcessingRequired {
				dividerStr := preparedArgFormat[1:]
				dividerVal, err := strconv.ParseFloat(dividerStr, 32)
				if err == nil {
					// 1. Convert arg to float
					val := (*item).(interface{})
					var floatVal float64
					switch val.(type) {
					case float64:
						floatVal = val.(float64)
					case int:
						floatVal = float64(val.(int))
					default:
						floatVal = 0
					}
					// 2. Divide arg / divider and multiply by 100
					percentage := (floatVal / dividerVal) * 100
					return strconv.FormatFloat(percentage, floatFormat, 2, 64)
				}
			}

		default:
			base = 10
		}
	}

	switch v := (*item).(type) {
	case string:
		argStr = v
	case int8:
		argStr = strconv.FormatInt(int64(v), base)
	case int16:
		argStr = strconv.FormatInt(int64(v), base)
	case int32:
		argStr = strconv.FormatInt(int64(v), base)
	case int64:
		argStr = strconv.FormatInt(v, base)
	case int:
		argStr = strconv.FormatInt(int64(v), base)
	case uint8:
		argStr = strconv.FormatUint(uint64(v), base)
	case uint16:
		argStr = strconv.FormatUint(uint64(v), base)
	case uint32:
		argStr = strconv.FormatUint(uint64(v), base)
	case uint64:
		argStr = strconv.FormatUint(v, base)
	case uint:
		argStr = strconv.FormatUint(uint64(v), base)
	case bool:
		argStr = strconv.FormatBool(v)
	case float32:
		argStr = strconv.FormatFloat(float64(v), floatFormat, precision, 32)
	case float64:
		argStr = strconv.FormatFloat(v, floatFormat, precision, 64)
	default:
		argStr = fmt.Sprintf("%v", v)
	}

	if !postProcessingRequired {
		return argStr
	}

	// 1. If integer numbers add filling
	if intNumberFormat {
		symbolsStr := preparedArgFormat[1:]
		symbolsStrVal, err := strconv.Atoi(symbolsStr)
		if err == nil {
			symbolsToAdd := symbolsStrVal - len(argStr)
			if symbolsToAdd > 0 {
				advArgStr := &strings.Builder{}
				advArgStr.Grow(len(argStr) + symbolsToAdd)

				for i := 0; i < symbolsToAdd; i++ {
					advArgStr.WriteByte('0')
				}
				advArgStr.WriteString(argStr)
				return advArgStr.String()
			}
		}
	}

	if floatNumberFormat && precision > 0 {
		pointIndex := strings.Index(argStr, ".")
		if pointIndex > 0 {
			advArgStr := &strings.Builder{}
			advArgStr.Grow(len(argStr) + precision)
			advArgStr.WriteString(argStr)
			numberOfSymbolsAfterPoint := len(argStr) - (pointIndex + 1)
			for i := numberOfSymbolsAfterPoint; i < precision; i++ {
				advArgStr.WriteByte(0)
			}
			return advArgStr.String()
		}
	}

	return argStr
}
