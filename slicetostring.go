package stringFormatter

import "strings"

// SliceToString function that converts slice of any type items to string in format {item}{sep}{item}...
// TODO(UMV): probably add one more param to wrap item in quotes if necessary
func SliceToString(data *[]any, separator *string) string {
	if len(*data) == 0 {
		return ""
	}

	sliceStr := &strings.Builder{}
	// init memory initially
	sliceStr.Grow(len(*data)*len(*separator)*2 + (len(*data)-1)*len(*separator))
	isFirst := true
	for _, item := range *data {
		if !isFirst {
			sliceStr.WriteString(*separator)
		}
		sliceStr.WriteString(Format("{0}", item))
		isFirst = false
	}

	return sliceStr.String()
}
