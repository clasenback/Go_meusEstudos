package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type meuSlice []int

func (r meuSlice) contemValor(n int) bool {
	for _, v := range r {
		if n == v {
			return true
		}
	}
	return false
}

func main() {
	var xi meuSlice
	var xs []string
	var n int
	for i := 0; i < 64; i++ {
		n = rand.Intn(512)
		if !xi.contemValor(n) {
			xi = append(xi, n)
			xs = append(xs, string(n))
		}
	}

	fmt.Println("===== ANTES DO SORT ====")
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("===== DEPOIS DO SORT ====")
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("==== APRESENTANDO VALORES EM PARES ====")
	for col := range xi {
		fmt.Printf("%v %v\t", xi[col], xs[col])
		if col%8 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Println()
	fmt.Println(len(xi))
}
