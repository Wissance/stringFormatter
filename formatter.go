package stringFormatter

import (
	"strconv"
	"strings"
)

/* Func that make string formatting from template
 * Template is a following "Hello {0}, It is time to go to the {1}, {0}!"
 * if you are passing to upper mentioned str template values John, library you get
 * following result: Hello John, It is time to go to the library, John!
 */
func Format(template string, args ...string)  (string, error){

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