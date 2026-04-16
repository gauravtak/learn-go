package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func changeValue(val int, val2 *int) {
	val = 0
	*val2 = 3
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// changing a parameter passed by value 
	// only changes the local copy not the original variable
	zeroVal(i)
	fmt.Println("zeroval:", i)

	// if you want to change the value of variable i
	// You need to use pointer
	// Assigning a value to a dereferenced pointer
	// changes the value at the referenced address
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	a, b := 0, 0

	fmt.Printf("a: %d and b: %d\n", a, b)
	changeValue(a, &b)
	// here a will not change but b will
	fmt.Printf("a: %d and b: %d\n", a, b)
}
