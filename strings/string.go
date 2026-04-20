package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello"

	// i is the index 
	// v is the unicode id (code-point)
	for i, v := range s {
		// fmt.Println(i)
		// fmt.Println(v)
		fmt.Printf("%#U is at %d\n", v, i)
	}
	
	// length of string 
	fmt.Println("length of string", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	fmt.Println("\nrune count in this string is: ", utf8.RuneCountInString(s))

	fmt.Println("hello this is end of this functions")
}