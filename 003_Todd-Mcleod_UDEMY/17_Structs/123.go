package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

func main() {
	eu := pessoa{"Alex", "Back", 44, 1.75}
	ela := pessoa{"Norma", "Sizilio", 41, 1.55}
	fmt.Printf("%v \t %#v \t %T \n", eu, eu, eu)
	fmt.Printf("%v \t %#v \t %T \n", ela, ela, ela)
}
