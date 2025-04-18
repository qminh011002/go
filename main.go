package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func measure(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	var s Speaker
	fmt.Println(s)

	d := Dog{}
	measure(d)
	fmt.Println(s == d)
}
