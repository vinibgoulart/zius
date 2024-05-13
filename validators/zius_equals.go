package validators

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/vinibgoulart/zius/helpers"
)

type EqualsValidator struct {
	BaseValidator
}

func (v *EqualsValidator) Validate(value interface{}) error {
	message := MessageGet(v.TagMessage, fmt.Sprintf("%s must be equals to %s", *v.StructField, *v.TagValue))

	regex := helpers.CommaToVerticalBar(v.TagValue)
	ok := regexp.MustCompile(regex).MatchString(value.(string))
	if !ok {
		return errors.New(message)
	}

	return nil
}
