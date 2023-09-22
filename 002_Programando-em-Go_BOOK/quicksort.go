package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))
	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido!", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}
	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}
	ns := make([]int, len(numeros))
	copy(ns, numeros)
	indicePivo := len(ns) / 2
	pivo := ns[indicePivo]
	ns = append(ns[:indicePivo], ns[indicePivo+1:]...)
	menores, maiores := particionar(ns, pivo)
	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}
	return menores, maiores
}
