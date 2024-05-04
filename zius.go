package zius

import (
	"fmt"
	"reflect"

	"github.com/vinibgoulart/zius/parser"
	"github.com/vinibgoulart/zius/validators"
)

func Validate(s interface{}) error {
	structType := reflect.TypeOf(s)

	if structType.Kind() != reflect.Struct {
		panic("zius requires a struct")
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		zius := field.Tag.Get("zius")

		tags := parser.SplitCommaIgnoringTagValue(&zius)

		for _, tag := range tags {
			parsedTag := parser.ParseTag(&tag)

			validator, ok := validators.GetValidator(parsedTag, &field.Name)
			if !ok {
				return fmt.Errorf("unknown validator: %s", *parsedTag.TagType)
			}

			value := reflect.ValueOf(s).Field(i).Interface()
			err := validator.Validate(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
