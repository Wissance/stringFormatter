package stringFormatter

import (
	"unicode"
)

type FormattingStyle string

const (
	NoFormatting FormattingStyle = "no_formatting"
	Camel        FormattingStyle = "camel"
	Snake        FormattingStyle = "snake"
	Kebab        FormattingStyle = "kebab"
)

type styleStats struct {
	SignIndexes []int
}

var styles = map[rune]FormattingStyle{
	'_': Snake,
	'-': Kebab,
}

func SetStyle(text *string, style FormattingStyle) string {
	// 1. Slice text by spaces and prepare every item separately
	// 2. Find style marker (_, -, or transition lowerUpper)
	// 3. Convert
	return ""
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
func defineFormattingStyle(text *string) map[FormattingStyle]styleStats {
	// symbol analyze, for camel case pattern -> aA, for kebab -> a-a, for snake -> a_a
	var stats map[FormattingStyle]styleStats = map[FormattingStyle]styleStats{
		Kebab: styleStats{},
		Snake: styleStats{},
		Camel: styleStats{},
	}
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
		if style != NoFormatting {
			// unfortunately copy and create new value in the map :(
			newStats := append(stats[style].SignIndexes, pos)
			stats[style] = styleStats{SignIndexes: newStats}
		}
	}
	return stats
}

func isUpper(symbol rune) bool {
	return unicode.IsUpper(symbol) && unicode.IsLetter(symbol)
}
