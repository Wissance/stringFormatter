package stringFormatter

import (
	"strconv"
	"strings"
)

/* Func that make string formatting from template
 * Template is a following "Hello {0}, It is time to go to the {1}, {0}!
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

