package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	output := input * 2
	fmt.Println(output)

	var lista []int
	lista = append(lista, input, output)
	fmt.Println(lista)

	mapa := make(map[string]int)
	mapa["input"] = input
	mapa["output"] = output
	for k, v := range mapa {
		fmt.Println(k, v)
	}
}
