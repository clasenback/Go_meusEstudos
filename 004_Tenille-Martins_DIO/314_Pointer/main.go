package main

import "fmt"

func inicial(meuPonteiro *int) {
	*meuPonteiro = 0
}

func main() {
	x := 5
	inicial(&x)
	fmt.Println(x)
}
