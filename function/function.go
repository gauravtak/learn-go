package main

import "fmt"

// Function with single return value
func sum(a int, b int) int {
	return a + b
}

// Function with multiple return value
func vals() (int, int) {
	return 3, 7
}

// Variadic functions
// can take any number of arguments

func sumOfNaturalNumbers(nums ...int) {
	// Type of nums is []int
	// Can interate over it
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Anonymous functions
var k func() int

// Closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion
func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	fmt.Println("This is a function")
	var a int
	a = sum(3, 4)
	fmt.Println("sum of two numbers", a)

	c, d := vals()
	fmt.Println("Multi value returning function: ", c, d)
	// calling variadic function
	sumOfNaturalNumbers(1, 2, 3, 4)

	nat := []int{1, 2, 3, 4, 5}
	sumOfNaturalNumbers(nat...)

	e := 0
	f := 0
	if 10 > 5 {
		k = func() int {
			e = 5
			return e
		}
	} else {
		k = func() int {
			e = 10
			return e
		}
	}

	fmt.Println("e is: ", e)
	f = k()
	fmt.Println("f is: ", f)
	fmt.Println("e is: ", e)

	// calling closure
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())

	// calling recursion
	x := fact(0)
	y := fact(5)
	fmt.Println("x is: ", x)
	fmt.Println("y is: ", y)

}
