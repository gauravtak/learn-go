package main

import "fmt"

func main() {
	if num := 15; num < 12 {
		fmt.Println("Num is less than 12")
	} else if num <= 0 {
		fmt.Println("Num is less than 0")
	} else {
		fmt.Println("Num")
	}
}
