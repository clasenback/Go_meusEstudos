package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	estrut1 := pessoa{"Alex", 44}
	estrut2 := pessoa{"Norma", 41}
	fmt.Println(estrut1)
	fmt.Println(estrut2)
	fmt.Printf("%T \t %T \t %T \n", estrut1, estrut1.nome, estrut1.idade)
	fmt.Printf("%v \t %v \t %v \n", estrut1, estrut1.nome, estrut1.idade)
}
