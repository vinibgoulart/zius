package validators

import (
	"errors"
	"fmt"
	"regexp"
)

type RegexValidator struct {
	BaseValidator
}

func (v *RegexValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s must be a valid regex %s", *v.Field, *v.TagValue))

	ok := regexp.MustCompile(*v.TagValue).MatchString(value.(string))
	if !ok {
		return errors.New(message)
	}

	return nil
}
