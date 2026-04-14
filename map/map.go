package main

import (
	"fmt"
	"log"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 2
	m["k2"] = 3
	m["k3"] = 8

	fmt.Println("map", m)


	if err, ok := m["k3"]; ok {
		fmt.Println("This key exists")
	} else {
		log.Fatal("What is this", err)
	}

}
