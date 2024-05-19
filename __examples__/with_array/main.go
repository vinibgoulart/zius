package main

import (
	zius "github.com/vinibgoulart/zius"
)

type WithArray struct {
	Name           string   `zius:"required,string:Name must be a string"`
	ChildrensNames []string `zius:"required,array:ChildrensNames must be an array of string"`
}

func main() {
	ws := &WithArray{
		Name: "Vinicius",
		ChildrensNames: []string{
			"Vinicius Daughter",
			"Vinicius Son",
		},
	}

	_, err := zius.Validate(*ws)

	if err != nil {
		panic(err)
	} else {
		println("Valid!")
	}
}
