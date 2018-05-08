package main

import (
	"fmt"
)

type Try struct {
	User string `json:"user"`
	Sucker string `json:"Sucker"`
}

type TryFail struct {
	User string `json:"user"`
}


func AnotherTest() (*Try) {
	return &Try{"hello?", "how are you"}
}

func main() {
	fmt.Println(AnotherTest() == nil)
}