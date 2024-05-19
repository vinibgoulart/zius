package main

import (
	zius "github.com/vinibgoulart/zius"
)

type WithStructBody struct {
	Color string `zius:"required"`
}

type WithStructClothes struct {
	Body WithStructBody `zius:"required,struct={WithStructBody}:Body is required"`
}

type WithStructAddress struct {
	Street string `zius:"required"`
	Number int    `zius:"required"`
}

type WithStruct struct {
	Name    string            `zius:"required,string:Name must be a string"`
	Address WithStructAddress `zius:"required,struct:Address is required"`
	Clothes WithStructClothes `zius:"required,struct:Clothes is required"`
}

func main() {
	ws := &WithStruct{
		Name: "Vinicius",
		Address: WithStructAddress{
			Street: "Main Street",
			Number: 123,
		},
		Clothes: WithStructClothes{
			Body: WithStructBody{
				Color: "Blue",
			},
		},
	}

	_, err := zius.Validate(*ws)

	if err != nil {
		panic(err)
	} else {
		println("Valid!")
	}
}
