package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sorvete   []string
}

func main() {
	// EXERCÍCIO 127: STRUCT WITH SLICE
	fmt.Println("----------------")
	fmt.Println("EX. 127: STRUCT WITH SLICE")
	eu := pessoa{"Alex", "Back", []string{"passas ao rum", "limão", "creme"}}
	norma := pessoa{"Norma", "Sizilio", []string{"chocolate", "cereja"}}
	pessoas := []pessoa{eu, norma}
	for _, p := range pessoas {
		fmt.Println(p.nome)
		for _, f := range p.sorvete {
			fmt.Printf("%T \t %v \n", f, f)
		}
	}

	// EXERCÍCIO 128: MAP STRUCT
	fmt.Println("----------------")
	fmt.Println("EX. 128: MAP STRUCT")
	registros := make(map[string]pessoa)
	for _, p := range pessoas {
		registros[p.sobrenome] = p
	}
	for k, v := range registros {
		fmt.Println(k)
		for _, v := range v.sorvete {
			fmt.Println(v)
		}
	}
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
	● first name
	● last name
	● favorite ice cream flavors
Create two VALUES of TYPE person.
Print out the values, ranging over the elements in the slice which stores the favorite flavors.
https://go.dev/play/p/nLcea3bIb7h
*/

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.

*/
