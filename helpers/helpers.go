package helpers

import (
	"regexp"
	"strings"
)

func BracketsRemove(s string) string {
	return s[1 : len(s)-1]
}

func CommaToVerticalBar(valid *string) string {
	parts := strings.Split(*valid, ",")
	for i, part := range parts {
		parts[i] = regexp.QuoteMeta(part)
	}
	return strings.Join(parts, "|")
}

func CommaIgnoringTagValueSplit(input *string) []string {
	var result []string
	var current string
	var bracketDepth int

	for _, char := range *input {
		switch char {
		case ',':
			if bracketDepth == 0 {
				result = append(result, current)
				current = ""
			} else {
				current += string(char)
			}
		case '{':
			bracketDepth++
			current += string(char)
		case '}':
			bracketDepth--
			current += string(char)
		default:
			current += string(char)
		}
	}
	if current != "" {
		result = append(result, current)
	}

	return result
}
