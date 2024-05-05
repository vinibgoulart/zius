package validators

import (
	"errors"
	"fmt"
)

type IntegerValidator struct {
	BaseValidator
}

func (v *IntegerValidator) Validate(value interface{}) error {
	message := GetMessage(v.TagMessage, fmt.Sprintf("%s must be an integer", *v.StructField))

	if _, ok := value.(int); !ok {
		return errors.New(message)
	}

	return nil
}
