package main

import "fmt"

var middleName string = "zxc"

func main() {
	// var age int = 30
	// var name string = "John"
	// var name1 = "Jane"

	// var r rune

	// fmt.Println(r)
	// count := 10

	// Default values
	// Numeric Type: 0
	// Boolean Type: false
	// String Type: ""
	// Pointer, Slices, Maps, Functions, and Structs: nil

	// ---- SCOPE

	zero := 0b0 + 0b1 + 0b1
	fmt.Println(zero)
}

func printName() {
	firstName := "Michael"

	fmt.Println(firstName)
}
