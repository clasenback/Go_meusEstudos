package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40

	// Statement; Statement
	// é uma boa prática reduzir ao máximo o escopo de uma variável
	// então, se a variável será usada apenas dentro de um IF ou FOR
	// uma variável temporária pode ser empregada desta forma.
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z é %v and é MAIOR OU IGUAL a %v \n", z, x)
	} else {
		fmt.Printf("z é %v and é MENOR a %v \n", z, x)
	}
}
