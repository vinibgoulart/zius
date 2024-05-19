package main

import (
	zius "github.com/vinibgoulart/zius"
)

type Simple struct {
	Name   string `zius:"required:Name is required,string"`
	Age    int    `zius:"required,int"`
	Gender string `zius:"required,string,equals={BOY,GIRL}:Should be BOY or GIRL"`
}

func main() {
	s := &Simple{
		Name:   "John",
		Age:    30,
		Gender: "BOY",
	}

	_, err := zius.Validate(*s)

	if err != nil {
		panic(err)
	} else {
		println("Valid!")
	}
}
