package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("PinPan \t")
		} else if i%3 == 0 {
			fmt.Print("Pin \t")
		} else if i%5 == 0 {
			fmt.Print("Pan \t")
		} else {
			fmt.Print(i, "\t")
		}
	}
}

/*
Imprimir números de 1 a 100.
Quando múltiplo de 3, imprimir "Pin" no lugar do número.
Quando múltiplo de 5, imprimir "Pan" no lugar do número.
Se múltiplo de 3 e 5, imprimir "Pin"
*/
