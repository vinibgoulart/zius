package validators

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/vinibgoulart/zius/parser"
)

type EqualsValidator struct {
	BaseValidator
}

func (v *EqualsValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s must be equals to %s", *v.StructField, *v.TagValue))

	regex := parser.CommaToVerticalBar(v.TagValue)
	ok := regexp.MustCompile(regex).MatchString(value.(string))
	if !ok {
		return errors.New(message)
	}

	return nil
}
