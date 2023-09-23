package main

import "fmt"

func main() {

	// var y string
	var x int = 8
	n, err := fmt.Scan(x)

	// SWITCH SIMPLES
	// fallthrough força que o bloco do próximo 'case' seja EXECUTADO,
	// sem avaliar o 'case' associado
	switch {
	case x%2 == 0:
		fmt.Println(x, "é par.", x%2 == 0)
		fallthrough
	case x%3 == 0:
		fmt.Println(x, "é múltiplo de 3.", x%3 == 0)
		fallthrough
	case x%5 == 0:
		fmt.Println(x, "é múltiplo de 5.", x%5 == 0)
	default:
		fmt.Println(x, "não tem fatores menores que 5.")
	}

	//SWITCH X
	switch x {
	case 2:
		fmt.Println("Dois")
		fallthrough
	case 3:
		fmt.Println("Três")
		fallthrough
	case 5:
		fmt.Println("Cinco")
	default:
		fmt.Println("Não é 2, 3 ou 5.")
	}

	fmt.Println(x, "\t", n, "\t", err)
}
