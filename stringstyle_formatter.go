package stringFormatter

import (
	"strings"
	"unicode"
)

type FormattingStyle string

const (
	Camel FormattingStyle = "camel"
	Snake FormattingStyle = "snake"
	Kebab FormattingStyle = "kebab"
)

type styleInc struct {
	Index int
	Style FormattingStyle
}

var styles = map[rune]FormattingStyle{
	'_': Snake,
	'-': Kebab,
}

func SetStyle(text *string, style FormattingStyle) string {
	if text == nil {
		return ""
	}
	sb := strings.Builder{}
	sb.Grow(len(*text))
	stats := defineFormattingStyle(text)
	startIndex := 0
	endIndex := 0
	// todo UMV think how to process ....
	// we could have many stats at the same time, probably we should use some config in the future
	// iterate over the map
	for _, v := range stats {
		endIndex = v.Index
		sb.WriteString((*text)[startIndex:endIndex])
		startIndex = v.Index
		if style == v.Style {
			sb.WriteString((*text)[:endIndex])
		} else {
			switch style {
			case Kebab:
				sb.WriteString("-")
				break
			case Snake:
				sb.WriteString("_")
				break
			case Camel:
				// in case of convert to Camel we should skip v.Index (because it is _ or -)
				sb.WriteRune(unicode.ToUpper(rune((*text)[endIndex+1])))
				startIndex += 1
				break
			}
			startIndex += 1
		}

	}
	sb.WriteString((*text)[startIndex:])
	if style != Camel {
		return strings.ToLower(sb.String())
	}
	return sb.String()
}

// defineFormattingStyle
/* This function defines what formatting style is using in text
 * If there are no transitions between symbols then here we have NoFormatting style
 * Didn't decide yet what to do if we are having multiple signatures
 * i.e. multiple_signs-at-sameTime .
 * Parameters:
 *    - text - a sequence of symbols to check
 * Returns: formatting style using in the text
 */
func defineFormattingStyle(text *string) []styleInc {
	// symbol analyze, for camel case pattern -> aA, for kebab -> a-a, for snake -> a_a
	inclusions := make([]styleInc, 0)
	runes := []rune(*text)
	for pos, char := range runes {
		// define style and add stats
		style, ok := styles[char]
		if !ok {
			// 1. Probably current symbol is not a sign and we should continue
			if pos > 0 && pos < len(runes)-1 {
				charIsUpperCase := isUpper(char)
				prevChar := runes[pos-1]
				prevCharIsUpperCase := unicode.IsUpper(prevChar)
				if charIsUpperCase != prevCharIsUpperCase {
					style = Camel
				}
			}
		}
		if style != "" {
			inclusions = append(inclusions, styleInc{Index: pos, Style: style})
		}
	}
	return inclusions
}

func isUpper(symbol rune) bool {
	return unicode.IsUpper(symbol) && unicode.IsLetter(symbol)
}
