package stringFormatter

import "strings"

type MapLineFormat string

const (
	keyName   = "key"
	keyArg    = "{" + keyName + "}"
	valueName = "value"
	valueArg  = "{" + valueName + "}"
)

const (
	KeyValueWithArrowSepFormat     MapLineFormat = keyArg + " => " + valueArg
	KeyValueWithSemicolonSepFormat MapLineFormat = keyArg + " : " + valueArg
	ValueOnly                      MapLineFormat = valueArg
)

func MapToString[TK string | int | uint | int32 | int64 | uint32 | uint64, TV any](data *map[TK]TV, format MapLineFormat, lineSeparator string) string {

	if data == nil || len(*data) == 0 {
		return ""
	}
	var mapStr = &strings.Builder{}
	empty := true
	mapStr.Grow(len(*data) * 50)
	lineData := map[string]interface{}{}
	for k, v := range *data {
		lineData[keyName] = k
		lineData[valueName] = v
		line := FormatComplex(string(format), lineData)
		// append
		if !empty {
			mapStr.WriteString(lineSeparator)
		}
		mapStr.WriteString(line)
		empty = false
	}
	return mapStr.String()
}
