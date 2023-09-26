package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

// Embedded struct
type agenteSecreto struct {
	identidadeReal pessoa
	codinome       string
	licencaMatar   bool
	nome           string
}

func main() {
	// Basic struct
	fmt.Println("----------------")
	fmt.Println("BASIC STRUCTS")
	eu := pessoa{"Alex", "Back", 44, 1.75}
	ela := pessoa{"Norma", "Sizilio", 41, 1.55}
	fmt.Printf("valor: \t\t%v \nconteúdo: \t%#v \ntipo: \t\t%T \npointer: \t%v \n", eu, eu, eu, &eu)
	fmt.Printf("valor: \t\t%v \nconteúdo: \t%#v \ntipo: \t\t%T \npointer: \t%v \n", ela, ela, ela, &ela)

	// Embedded struct
	fmt.Println("----------------")
	fmt.Println("EMBEDDED STRUCT")
	ag := agenteSecreto{eu, "Xandeca", true, "Bond"}
	fmt.Printf("valor: \t%v \nconteúdo: \t%#v \ntipo: \t%T \n", ag, ag, ag)
	fmt.Println(ag.codinome, ag.identidadeReal.nome, ag.nome)

	// Anonymous struct
	fmt.Println("----------------")
	fmt.Println("ANONYMOUS STRUCT")
	mae := struct {
		nome      string
		sobrenome string
		idade     int
		altura    float64
	}{"Maria de Lourdes", "Clasen Back", 72, 1.60}
	fmt.Printf("tipo: \t\t%T \nconteúdo: \t%v \npointer: \t%v\n", mae, mae, &mae)
	fmt.Printf("pointer nome: \t\t%v \t%v \t%T\n", mae.nome, &mae.nome, mae.nome)
	fmt.Printf("pointer sobrenome: \t%v \t\t%v \t%T\n", mae.sobrenome, &mae.sobrenome, mae.sobrenome)
	fmt.Printf("pointer idade: \t\t%v \t\t\t%v \t%T\n", mae.idade, &mae.idade, mae.idade)
	fmt.Printf("pointer altura: \t%v \t\t\t%v \t%T\n", mae.altura, &mae.altura, mae.altura)
}
