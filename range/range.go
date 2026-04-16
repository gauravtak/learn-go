package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5,6}

	sum := 0

	// the range provides both index and value
	// for each entry, here we do not need index
	// so we just used _ (blank identifier)
	for _, num := range nums {
		sum += num
	}

	// here I have also used the index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	fmt.Println(nums)

	// using range to iterate over maps
	fmt.Println("Iterating over maps")
	fruits := map[string]string{"a": "apple", "b": "banana"}

	for k := range fruits {
		fmt.Println("key: ", k)
	}

}