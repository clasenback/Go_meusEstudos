package main

import (
	"fmt"
	"sort"
)

func main() {

	quadrados := make(map[int]int)
	for i := 0; i <= 10; i++ {
		quadrados[i] = i * i
	}

	fmt.Println("-+- Print do map direto -+-")
	for k, v := range quadrados {
		fmt.Printf("%v \t %v \n", k, v)
	}
	fmt.Println("-+- Print do map após sorted key slice -+-")

	numeros := []int{}
	for i, _ := range quadrados {
		numeros = append(numeros, i)
	}

	sort.Ints(numeros)

	for _, k := range numeros {
		fmt.Printf("%v \t %v \n", k, quadrados[k])
	}
}
