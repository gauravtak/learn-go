package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	// GO is garbage collected so you can safely return a pointer to the local variable
	// Here the local variable is p
	p := person{name: name}
	p.age = 32
	return &p
}

func main() {
	p := person{name: "Alice"}
	fmt.Println("p is: ", p)

	fmt.Println(person{name: "Kat", age: 20})
	// Using our constructor function
	p2 := newPerson("John")
	fmt.Println("p2 is: ", p2)

	//Accessing fields of a struct
	s := person{name: "Ajay", age: 18}
	fmt.Println("age of s: ", s.age)
	fmt.Println("name of s: ", s.name)

	sp := &s
	fmt.Println(sp.age) // the pointer will get dereferenced automatically

	// If you only have a single value for struct type then...
	dog := struct {
		name  string
		breed string
	}{
		"Rony",
		"German Shepherd",
	}

	fmt.Println(dog)
}
