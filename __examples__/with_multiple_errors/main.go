package main

import (
	"fmt"

	zius "github.com/vinibgoulart/zius"
)

type WithMultipleErrors struct {
	Name   string `zius:"required,string:Name must be a string"`
	Age    int    `zius:"required,int"`
	Gender string `zius:"required,string,equals={BOY,GIRL}:Should be BOY or GIRL"`
}

func main() {
	wme := &WithMultipleErrors{
		Name:   "",
		Age:    30,
		Gender: "HELLO WORLD",
	}

	errors, _ := zius.Validate(*wme)

	if len(errors) > 0 {
		panic(fmt.Sprint(errors))
	} else {
		println("Valid!")
	}
}
