package main

import "fmt"

func main() {
	// fmt.Println("Declaring initial variable outside")
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// fmt.Println("Declaring initial variable inside")
	// for j := 0; j < 5; j++ {
	// 	fmt.Println(j)
	// }

	fmt.Println("Declaring initial variable with range")
	for k := range 10 {
		if k%2 == 0 {
			fmt.Println("Gangster")
		}
		fmt.Println(k)
	}
}
