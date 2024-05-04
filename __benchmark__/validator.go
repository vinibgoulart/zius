package main

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	zius "github.com/vinibgoulart/zius"
)

type Person struct {
	Name  string `zius:"required:Name is required" validate:"required"`
	Age   int    `zius:"required:Age is required" validate:"required"`
	Email string `zius:"required:Email is required" validate:"required,email"`
}

func main() {
	p := &Person{
		Name:  "Vinicius",
		Age:   20,
		Email: "viblaziusgoulart@gmail.com",
	}

	startValidator := time.Now()
	validate := validator.New()
	validate.Struct(*p)
	endValidator := time.Now()

	fmt.Println("Validator -", endValidator.Sub(startValidator))

	start := time.Now()
	zius.Validate(*p)
	end := time.Now()

	fmt.Println("Zius -", end.Sub(start))
}
