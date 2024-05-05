package validators

import (
	"github.com/vinibgoulart/zius/parser"
)

type Validator interface {
	Validate(value interface{}) error
}

type BaseValidator struct {
	StructField *string
	TagValue    *string
	TagMessage  *string
}

func GetValidator(pt *parser.ParsedTag, structField *string, structType *interface{}) (Validator, bool) {
	baseValidator := BaseValidator{TagMessage: pt.TagMessage, TagValue: pt.TagValue, StructField: structField}

	switch *pt.TagType {
	case RequiredTag:
		return &RequiredValidator{baseValidator}, true
	case IntegerTag:
		return &IntegerValidator{baseValidator}, true
	case StringTag:
		return &StringValidator{baseValidator}, true
	case NumberTag:
		return &NumberValidator{baseValidator}, true
	case EmailTag:
		return &EmailValidator{baseValidator}, true
	case PhoneTag:
		return &PhoneValidator{baseValidator}, true
	case RegexTag:
		return &RegexValidator{baseValidator}, true
	case EqualsTag:
		return &EqualsValidator{baseValidator}, true
	default:
		return nil, false
	}
}
