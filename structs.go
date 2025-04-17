package main

import "fmt"

// Structs are defined using the 'type' and 'struct' keywords followed by
// curly braces '{}'
// Fields are defined with a name and a type

func main() {
	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       23,
	}

	fmt.Println(p)
}
