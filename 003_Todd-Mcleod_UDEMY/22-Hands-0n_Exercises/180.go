package main

import "fmt"

func main() {
	eu := pessoa{"Alex", 44}
	fmt.Println(eu)
	eu = mudaNome(eu, "Alexander")
	fmt.Println(eu)
	mudaNomePointer(&eu, "Xande")
	fmt.Println(eu)
}

type pessoa struct {
	nome  string
	idade int
}

func mudaNome(p pessoa, s string) pessoa {
	p.nome = s
	return p
}

func mudaNomePointer(p *pessoa, s string) {
	p.nome = s
}

/*
Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
	○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS
	○ this function will return nothing
● don't use methods
*/
