package stringFormatter

import (
	"strings"
	"unicode"
)

type FormattingStyle string
type CaseSetting int

const (
	Camel FormattingStyle = "camel"
	Snake FormattingStyle = "snake"
	Kebab FormattingStyle = "kebab"
)

const (
	ToUpper   CaseSetting = 1
	ToLower   CaseSetting = 2
	NoChanges CaseSetting = 3
)

type styleInc struct {
	Index int
	Style FormattingStyle
}

var styleSigns = map[rune]FormattingStyle{
	'_': Snake,
	'-': Kebab,
}

// SetStyle is a function that converts text with code to defined code style.
/* Set text like a code style to on from FormattingStyle (Camel, Snake, or Kebab)
 * conversion of abbreviations like JSON, USB, and so on is going like a regular text
 * for current version, therefore they these abbreviations could be in a different
 * case after conversion.
 * Case settings apply in the following order : 1 - textCase, 2 - firstSymbol.
 * If you are not applying textCase to text converting from Camel to Snake or Kebab
 * result is lower case styled text. textCase does not apply to Camel style.
 * Parameters:
 * - text - pointer to text
 * - style - new code style
 * - firstSymbol - case settings for first symbol
 * - textCase - case settings for whole text except first symbol
 * Returns : new string with formatted line
 */
func SetStyle(text *string, style FormattingStyle, firstSymbol CaseSetting, textCase CaseSetting) string {
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
		if endIndex < startIndex {
			continue
		}
		sb.WriteString((*text)[startIndex:endIndex])
		startIndex = v.Index

		switch style {
		case Kebab:
			sb.WriteString("-")
			break
		case Snake:
			sb.WriteString("_")
			break
		case Camel:
			// in case of convert to Camel we should skip v.Index (because it is _ or -)
			if v.Style == Camel {
				sb.WriteRune(unicode.ToUpper(rune((*text)[endIndex])))
			} else {
				sb.WriteRune(unicode.ToUpper(rune((*text)[endIndex+1])))
			}
			startIndex += 1
			break
		}
		if v.Style != Camel {
			startIndex += 1
		}
	}
	sb.WriteString((*text)[startIndex:])
	result := strings.Builder{}
	if style != Camel {
		switch textCase {
		case ToUpper:
			result.WriteString(strings.ToUpper(sb.String()[1:]))
			break
		case ToLower:
			result.WriteString(strings.ToLower(sb.String()[1:]))
			break
		case NoChanges:
			result.WriteString(sb.String()[1:])
			break
		}
	} else {
		result.WriteString(sb.String()[1:])
	}

	switch firstSymbol {
	case ToUpper:
		return strings.ToUpper(sb.String()[:1]) + result.String()
	case ToLower:
		return strings.ToLower(sb.String()[:1]) + result.String()
	case NoChanges:
		return sb.String()[:1] + result.String()
	}
	return ""
}

func GetFormattingStyleOptions(style string) (FormattingStyle, CaseSetting, CaseSetting) {
	styleLower := strings.ToLower(style)
	var formattingStyle FormattingStyle
	firstSymbolCase := ToLower
	textCase := NoChanges
	switch styleLower {
	case string(Camel):
		formattingStyle = Camel
		break
	case string(Snake):
		formattingStyle = Snake
		break
	case string(Kebab):
		formattingStyle = Kebab
		break
	}

	runes := []rune(style)
	firstSymbolIsUpper := isSymbolIsUpper(runes[0])
	if firstSymbolIsUpper {
		firstSymbolCase = ToUpper
	}

	if formattingStyle != Camel {

	}

	return formattingStyle, firstSymbolCase, textCase
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
		style, ok := styleSigns[char]
		if !ok {
			// 1. Probably current symbol is not a sign and we should continue
			if pos > 0 && pos < len(runes)-1 {
				charIsUpperCase := isSymbolIsUpper(char)
				prevChar := runes[pos-1]
				prevCharIsUpperCase := unicode.IsUpper(prevChar)
				if charIsUpperCase && !prevCharIsUpperCase {
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

func isSymbolIsUpper(symbol rune) bool {
	return unicode.IsUpper(symbol) && unicode.IsLetter(symbol)
}

func isStringIsUpper(str []rune) bool {
	isUpper := true
	for _, r := range str {
		if unicode.IsLetter(r) {
			isUpper = isUpper && unicode.IsUpper(r)
		}
	}
	return isUpper
}
