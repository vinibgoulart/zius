package validators

import (
	"errors"
	"fmt"
	"regexp"
)

type StringValidator struct {
	BaseValidator
}

func (v *StringValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s must be a string", *v.Field))

	regex := GetRegexForTag(StringTag)
	ok := regexp.MustCompile(regex).MatchString(value.(string))
	if !ok {
		return errors.New(message)
	}

	return nil
}
