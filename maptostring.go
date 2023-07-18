package stringFormatter

import "strings"

const (
	// KeyKey placeholder will be formatted to map key
	KeyKey = "key"
	// KeyValue placeholder will be formatted to map value
	KeyValue = "value"
)

// MapToString - format map keys and values according to format, joining parts with separator.
// Format should contain key and value placeholders which will be used for formatting, e.g.
// "{key} : {value}", or "{value}", or "{key} => {value}".
// Parts order in resulting string is not guranteed.
func MapToString[
	K string | int | uint | int32 | int64 | uint32 | uint64,
	V any,
](data map[K]V, format string, separator string) string {
	if len(data) == 0 {
		return ""
	}

	mapStr := &strings.Builder{}
	// assuming format will be at most two times larger after formatting part,
	// plus exact number of bytes for separators
	mapStr.Grow(len(data)*len(format)*2 + (len(data)-1)*len(separator))

	isFirst := true
	for k, v := range data {
		if !isFirst {
			mapStr.WriteString(separator)
		}

		line := FormatComplex(string(format), map[string]any{
			KeyKey:   k,
			KeyValue: v,
		})
		mapStr.WriteString(line)
		isFirst = false
	}

	return mapStr.String()
}
