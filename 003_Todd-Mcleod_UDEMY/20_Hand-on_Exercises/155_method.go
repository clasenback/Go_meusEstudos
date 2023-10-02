package main

import "fmt"

func main() {
	eu := person{first: "Alex", age: 44}
	eu.speak()
}

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Meu nome é %s e tenho %v anos.", p.first, p.age)
}

/*
● Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ age
● attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
