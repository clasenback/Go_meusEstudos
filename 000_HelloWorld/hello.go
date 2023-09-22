package main

import "fmt"

func main() {
	var n int
	var err error
	n, err = fmt.Println("Hello Word")
	fmt.Println(n, "caracteres")
	fmt.Println(err)

	fmt.Printf("%b caracteres \n", n)
	fmt.Printf("%x caracteres \n", n)

	var nome string
	fmt.Println("Qual seu nome?")
	fmt.Scanf(nome)
	fmt.Println("Olá,", nome)
}
