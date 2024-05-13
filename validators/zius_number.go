package validators

import (
	"errors"
	"fmt"
)

type NumberValidator struct {
	BaseValidator
}

func (v *NumberValidator) Validate(value interface{}) error {
	message := MessageGet(v.TagMessage, fmt.Sprintf("%s must be a number", *v.StructField))

	if _, ok := value.(float64); !ok {
		return errors.New(message)
	}

	return nil
}
