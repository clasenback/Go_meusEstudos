package main

import "fmt"

/* o operador 'iota' funciona como um contador,
bastando usar para a primeira variável ou constante,
as demais seguirão o padrão.
Pode ser usado com o operador nulo "_" para
desprezar algum valor indesejado.
*/

const (
	_  = iota // c0 == 0
	c1        // c1 == 1
	c2        // c2 == 2
	c3        // c3 == 3
	c4
	c5
	c6
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6)
	fmt.Printf("KB \t %d \t %#b\n", 1<<(10*c1), 1<<(10*c1))
	fmt.Printf("MB \t %d \t %#b\n", 1<<(10*c2), 1<<(10*c2))
	fmt.Printf("GB \t %d \t %#b\n", 1<<(10*c3), 1<<(10*c3))
	fmt.Printf("TB \t %d \t %#b\n", 1<<(10*c4), 1<<(10*c4))
	fmt.Printf("PB \t %d \t %#b\n", 1<<(10*c5), 1<<(10*c5))
	fmt.Printf("EB \t %d \t %#b\n", 1<<(10*c6), 1<<(10*c6))
}
