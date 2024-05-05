package validators

import (
	"errors"
	"fmt"
	"regexp"
)

type EmailValidator struct {
	BaseValidator
}

func (v *EmailValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s must be a valid email", *v.StructField))

	regex := GetRegexForTag(EmailTag)
	ok := regexp.MustCompile(regex).MatchString(value.(string))
	if !ok {
		return errors.New(message)
	}

	return nil
}
