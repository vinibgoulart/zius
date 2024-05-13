package validators

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/vinibgoulart/zius/parser"
)

type ArrayValidator struct {
	BaseValidator
}

func (v *ArrayValidator) Validate(value interface{}) error {
	arrayType := (*v.StructType).String()[2:]

	message := MessageGet(v.TagMessage, fmt.Sprintf("%s must be an array of %s", *v.StructField, arrayType))

	pt := &parser.ParsedTag{
		TagType:    &arrayType,
		TagValue:   nil,
		TagMessage: &message,
	}

	validator, ok := GetValidator(pt, v.StructField, v.StructType)
	if !ok {
		return errors.New(message)
	}

	valueSlice := reflect.ValueOf(value)
	for i := 0; i < valueSlice.Len(); i++ {
		item := valueSlice.Index(i).Interface()

		err := validator.Validate(item)
		if err != nil {
			return errors.New(message)
		}
	}

	return nil
}
