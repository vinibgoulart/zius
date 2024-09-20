package _struct

import (
	"fmt"
	"reflect"

	"github.com/vinibgoulart/zius/helpers"
	"github.com/vinibgoulart/zius/parser"
	"github.com/vinibgoulart/zius/validators"
)

type Error struct {
	Struct string
	Field  string
	Error  error
}

var AllErrors = make([]Error, 0)

func StructMapAll(s interface{}) []interface{} {
	AllErrors = make([]Error, 0)

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
		_zius := field.Tag.Get("zius")

		tags := helpers.CommaIgnoringTagValueSplit(&_zius)

		for _, tag := range tags {
			parsedTag := parser.ParseTag(&tag)

			value := reflect.ValueOf(s).Field(i).Interface()

			fmt.Println(*parsedTag)
			fmt.Println(value)

			validator, ok := validators.GetValidator(parsedTag, &field.Name, &field.Type)
			if !ok {
				return fmt.Errorf("unknown validator: %s", *parsedTag.TagType)
			}

			if validator == nil {
				continue
			}

			err := validator.Validate(value)

			fmt.Println(AllErrors)

			if err != nil {
				AllErrors = append(AllErrors, Error{
					Struct: reflect.TypeOf(s).Name(),
					Field:  field.Name,
					Error:  err,
				})
			}
		}
	}

	return nil
}
