package main

import (
	zius "github.com/vinibgoulart/zius"
)

type WithRegex struct {
	Name  string `zius:"required:Name is required,string:Name must be a string"`
	Age   int    `zius:"required:Age is required,int:Age must be an integer"`
	Email string `zius:"required:Email is required,email:Email must be a valid email"`
	Phone string `zius:"phone:Phone must be a valid phone number"`
	Star  string `zius:"required,regex={^[a-zA-Z]\\d{1,2}$}:Start must start with a letter and have 1 or 2 digits"`
}

func main() {
	wr := &WithRegex{
		Name:  "John",
		Age:   30,
		Email: "viblaziusgoulart@gmail.com",
		Phone: "5548999721289",
		Star:  "S20",
	}

	err := zius.Validate(*wr)

	if err != nil {
		panic(err)
	} else {
		println("Valid!")
	}
}
