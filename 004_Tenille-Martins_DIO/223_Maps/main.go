package main

import "fmt"

func main() {
	meuMapa := make(map[string]int)
	meuMapa["Primeiro"] = 1
	fmt.Println(meuMapa["Primeiro"])
	fmt.Println(meuMapa)
	meuMapa["Segundo"] = 2
	fmt.Println(meuMapa)
}
