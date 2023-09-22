package main

import "fmt"

func main() {
	var nome string = "Alex"
	var idade int8 = 44
	var altura float32 = 1.75
	var casado bool = true

	fmt.Println("Seu nome é:", nome)
	fmt.Println("Sua idade é:", idade)
	fmt.Println("Sua altura é:", altura)
	fmt.Println(nome, "é casado?", casado)
}
