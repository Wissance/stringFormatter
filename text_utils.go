package stringFormatter

type MapLineFormat string

const (
	keyName   = "{key}"
	valueName = "{value}"
)

const (
	KeyValueWithArrowSepFormat     MapLineFormat = keyName + " => " + valueName
	KeyValueWithSemicolonSepFormat MapLineFormat = keyName + " : " + valueName
	ValueOnly                      MapLineFormat = valueName
)

func MapToString[T any](data *map[string]T, format MapLineFormat) string {

	if data == nil || len(*data) == 0 {
		return ""
	}

	for k, v := range *data {

	}
}
