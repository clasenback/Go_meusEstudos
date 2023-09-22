package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	fmt.Println("some shit")
	defer dia2()
	fmt.Println("some shit")
	fmt.Println("some shit")

	fmt.Println("some shit")
	dia1()
	fmt.Println("some shit")

	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico")
}
