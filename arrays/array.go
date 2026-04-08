package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = [5]int{2: 1, 4: 2}
	fmt.Println(a)
	fmt.Println(b)

	// length of an array
	fmt.Println(len(a))


	// multi-dimensional array
	var twoD = [2][4]int{
		{1,2,3,5},
		{3,4,6,7},
	}

	fmt.Println("2d: ", len(twoD))
}
