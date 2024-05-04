package parser

import (
	"regexp"
	"strings"
)

type ParsedTag struct {
	TagType    *string
	TagValue   *string
	TagMessage *string
}

func ParseTag(s *string) *ParsedTag {
	v := strings.Split(*s, ":")

	tag := v[0]

	tagMessage := ""
	if len(v) > 1 {
		tagMessage = v[1]
	}

	vt := strings.Split(tag, "=")

	tagType := vt[0]

	tagValue := ""
	if len(vt) > 1 {
		tagValue = RemoveBrackets(vt[1])
	}

	return &ParsedTag{
		TagType:    &tagType,
		TagValue:   &tagValue,
		TagMessage: &tagMessage,
	}
}

func RemoveBrackets(s string) string {
	return s[1 : len(s)-1]
}

func CommaToVerticalBar(valid *string) string {
	parts := strings.Split(*valid, ",")
	for i, part := range parts {
		parts[i] = regexp.QuoteMeta(part)
	}
	return strings.Join(parts, "|")
}

func SplitCommaIgnoringTagValue(input *string) []string {
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
