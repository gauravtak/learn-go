package main

import "fmt"

func main() {
	// The difference between a slice and an array is that
	// slices are only defined by the type of elements they hold (not by the number of elements)

	var s []string
	fmt.Println(s)

	s = make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(s[1])

	s = append(s, "d")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println(c)

	fmt.Println(c[2:4])
}
