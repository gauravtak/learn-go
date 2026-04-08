package main

import "fmt"

func main() {
	var b int = 10
	var c string = "abcd"
	fmt.Println("######################## VARIABLE DECLARATIONS ########################")
	fmt.Printf("%d\n", b)
	fmt.Printf("%s\n", c)
	fmt.Println("######################## VARIABLE DECLARATIONS WITHOUT ASSIGNMENT ########################")
	var a int
	var cx string
	var d bool
	fmt.Printf("%d\n", a)
	fmt.Printf("%s\n", cx)
	fmt.Printf("%t\n", d)
	static()
	fmt.Println(true && false)
	fmt.Println(true || false)

	// Declaring multiple variables of same type at once
	ax, bx, cy := 0, 0, 0
	fmt.Println(ax)
	fmt.Println(bx)
	fmt.Println(cy)
}

// static typing for go
func static() {
	var a string = "10 "
	var b string = "hello"
	var c string
	c = a + b
	fmt.Printf("%s\n", c)
}
