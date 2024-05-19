package zius

import (
	"reflect"

	_struct "github.com/vinibgoulart/zius/struct"
)

func Validate(s interface{}) ([]_struct.Error, error) {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		panic("zius requires a struct")
	}

	allStructs := _struct.StructMapAll(s)

	for _, s := range allStructs {
		err := _struct.StructValidate(s)
		if err != nil {
			return nil, err
		}
	}

	if len(_struct.AllErrors) > 0 {
		return _struct.AllErrors, _struct.AllErrors[0].Error
	}

	return nil, nil
}
