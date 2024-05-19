package main

import (
	zius "github.com/vinibgoulart/zius"
)

type WithMessage struct {
	Name string `zius:"required:Name is required,string:Name must be a string"`
	Age  int    `zius:"required:Age is required,int:Age must be an integer"`
}

func main() {
	wm := &WithMessage{
		Name: "John",
		Age:  30,
	}

	_, err := zius.Validate(*wm)

	if err != nil {
		panic(err)
	} else {
		println("Valid!")
	}
}
