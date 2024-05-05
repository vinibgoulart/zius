package validators

import (
	"errors"
	"fmt"
)

type RequiredValidator struct {
	BaseValidator
}

func (v *RequiredValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s is required", *v.StructField))

	if value == nil || value == "" {
		return errors.New(message)
	}

	return nil
}
