package main

import (
	"errors"
	"fmt"
)

// In Go, a function can return two values
// as you can see here (int, error)
// see how the error is the last return value
// (error, int) this is wrong conventionally not technically, TL;DR: Go allows it, but you look dumb
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Another example
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil 
}


// Sentinel Error: Predeclared error variable specifying specific error
// other than that, this is also a higher level error
// which is wrapped up into Errorf at line: 39
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("Can't boil water")

func makeTea(arg int) error {
	if arg == 2  {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	result, err := divide(3, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)	
		} else {
			fmt.Println("f worked:", r)
		}
	}
	fmt.Println("Result:", result)

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy some tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		} 
	}
	fmt.Print("Tea is ready")

}