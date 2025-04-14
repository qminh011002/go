package main

import (
	"fmt"
)

// Sequences of Byte - Uint8
// Sequences of bytes represent text
//
//	Immutable
//	Access by index return byte value

// Runes and characters
//	 Similarities
//		Representing Chracter
//		Storage Size
// 	 Different
//		Unicode Support
//		Type and Size
//		Encoding and Hlding
// Practical Consideration
// 		Internationalization
//		Portability
//		Efficiency

func main() {
	// message := "Hello, \nGo!" // Double Quotes
	// message1 := "Hello, \t Go!"
	// message2 := "Hello, \rGo!"
	// rawMessage := `Hello, \nGo!` // Backtick

	// fmt.Println(message)
	// fmt.Println(message1)
	// fmt.Println(message2)
	// fmt.Println(rawMessage)

	// fmt.Println("Range of rawMessage", len(rawMessage))

	// fmt.Println(message[0])

	// greeting := "Hello"
	// name := "Alice"
	// msg := greeting + name
	// fmt.Println(msg)

	// str1 := "apple"
	// str := "apple"
	// fmt.Println(str > str1)

	// str := "Hello, Go!"

	// for _, v := range str {
	// 	// fmt.Printf("Character at index %d is %c \n", i, v)
	// 	fmt.Printf("%x \n", v)
	// }

	// fmt.Println("Rune count:", utf8.RuneCountInString(str))
	// s := "á"
	// fmt.Println(len(s))                    //2
	// fmt.Println(utf8.RuneCountInString(s)) //1

	// fmt.Println()

	// var ch rune = 'は'

	// fmt.Printf("%T", ch)

	// const NIHONGO = "日本語"

	jhello := "こんにちは" // Hello in Japanese

	for _, runeVal := range jhello {
		fmt.Println(runeVal)
	}

	// for _, v := range NIHONGO {
	// 	fmt.Println(v)
	// }

	r := '😊'

	fmt.Printf("%v \n", r)
	fmt.Printf("%c \n", r)

}
