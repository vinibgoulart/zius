package zius

import (
	"reflect"

	_struct "github.com/vinibgoulart/zius/struct"
)

func Validate(s interface{}) error {
	if reflect.TypeOf(s).Kind() != reflect.Struct {
		panic("zius requires a struct")
	}

	allStructs := _struct.StructMapAll(s)

	for _, s := range allStructs {
		err := _struct.StructValidate(s)
		if err != nil {
			return err
		}
	}

	return nil
}
