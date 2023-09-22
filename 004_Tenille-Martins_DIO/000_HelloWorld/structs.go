package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int8
}

func main() {
	pessoa1 := Pessoa{"Alex", "Back", 44}
	pessoa2 := Pessoa{"Norma", "Sizilio", 41}
	fmt.Println(pessoa1, "\n", pessoa2)
}
