package _struct

import (
	"fmt"
	"reflect"

	"github.com/vinibgoulart/zius/helpers"
	"github.com/vinibgoulart/zius/parser"
	"github.com/vinibgoulart/zius/validators"
)

func StructMapAll(s interface{}) []interface{} {
	result := make([]interface{}, 0)

	result = append(result, s)

	for i := 0; i < reflect.TypeOf(s).NumField(); i++ {
		field := reflect.TypeOf(s).Field(i)

		if field.Type.Kind() == reflect.Struct {
			result = append(result, reflect.ValueOf(s).Field(i).Interface())
			result = append(result, StructMapAll(reflect.ValueOf(s).Field(i).Interface())...)
		}
	}

	return result
}

func StructValidate(s interface{}) error {
	for i := 0; i < reflect.TypeOf(s).NumField(); i++ {
		field := reflect.TypeOf(s).Field(i)
		zius := field.Tag.Get("zius")

		tags := helpers.CommaIgnoringTagValueSplit(&zius)

		for _, tag := range tags {
			parsedTag := parser.ParseTag(&tag)

			value := reflect.ValueOf(s).Field(i).Interface()

			validator, ok := validators.GetValidator(parsedTag, &field.Name, &field.Type)
			if !ok {
				return fmt.Errorf("unknown validator: %s", *parsedTag.TagType)
			}

			if validator == nil {
				continue
			}

			err := validator.Validate(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
