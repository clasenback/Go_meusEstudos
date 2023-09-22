package main

import "fmt"

func main() {
	for i := 0; i < 13; i++ {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println(i, "é múltiplo de 2 e 3.")
		} else if i%2 == 0 || i%3 == 0 {
			fmt.Println(i, "é múltiplo de 2 ou 3.")
		} else {
			fmt.Println(i, "não é múltiplo de 2 ou 3.")
		}
	}
}
