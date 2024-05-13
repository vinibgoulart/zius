package parser

import (
	"strings"

	"github.com/vinibgoulart/zius/helpers"
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
		tagValue = helpers.BracketsRemove(vt[1])
	}

	return &ParsedTag{
		TagType:    &tagType,
		TagValue:   &tagValue,
		TagMessage: &tagMessage,
	}
}
