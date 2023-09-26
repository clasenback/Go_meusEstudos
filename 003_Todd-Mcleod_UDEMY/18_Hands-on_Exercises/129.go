package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {
	v1 := vehicle{
		engine: engine{true},
		make:   "BYD",
		model:  "shit",
		doors:  2,
		color:  "marrom"}
	v2 := vehicle{engine{false}, "Peugeot", "2008", 4, "vermelho"}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.make)
	fmt.Println(v2.make)

}

/*
● Create a type engine struct, and include this field
	○ electric bool
● Create a type vehicle struct, and include these fields
	■ engine
	■ make
	■ model
	■ doors
	■ color
● Create two VALUES of TYPE vehicle
	○ use a composite literal
● Print out each of these values.
● Print out a single field from each of these values.

*/
