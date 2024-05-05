package validators

import (
	"errors"
	"fmt"
	"regexp"
)

type PhoneValidator struct {
	BaseValidator
}

func (v *PhoneValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s must be a valid phone number", *v.StructField))

	regex := GetRegexForTag(PhoneTag)
	ok := regexp.MustCompile(regex).MatchString(value.(string))
	if !ok {
		return errors.New(message)
	}

	return nil
}
