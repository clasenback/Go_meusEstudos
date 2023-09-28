package main

import (
	"fmt"
	"time"
)

func makeNextNumberGenerator() func() uint {
	var i uint = 2
	return func() (ret uint) {
		ret = i
		i++
		return
	}
}

func makeEvenGenerator() func() uint {
	var i uint = 2
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makePrimeGenerator() func() (uint, []uint) {
	// Variáveis que serão capturadas pela closure
	primos := []uint{2}
	i := primos[len(primos)-1]

	// Defina a closure que retorna um uint e um slice de uints
	return func() (uint, []uint) {
		ret := primos[len(primos)-1]
		if ret == 2 {
			i++
		} else {
			i += 2
		}
		isPrime := false
		for !isPrime {
			for _, primo := range primos {
				if i%primo == 0 {
					i += 2
					if i%5 == 0 {
						i += 2
					}
					break
				} else if primo == primos[len(primos)/3] {
					isPrime = !isPrime
				}
			}
		}
		primos = append(primos, i)
		return i, primos
	}
}

func addAllBelow(n uint, xi []uint) uint {
	/* 	adiciona os elementos de um slice de valores
	   	ordenados até o elemento com valor n */
	var total uint
	for _, v := range xi {
		if v < n {
			total += v
		} else {
			break
		}
	}
	return total
}

func addAll(xi []int) int {
	/* 	sum is a function that takes a slice of numbers
	   	and adds them together. What would its function
	   	signature look like in Go?
	*/
	var total int
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {

	vezes := 10

	fmt.Print("Dez números: \t")
	next := makeNextNumberGenerator()
	for i := 0; i < vezes; i++ {
		fmt.Printf("\t%v", next())
	}
	fmt.Println("")

	fmt.Print("Dez pares: \t")
	next = makeEvenGenerator()
	for i := 0; i < vezes; i++ {
		fmt.Printf("\t%v", next())
	}
	fmt.Println("")

	fmt.Print("Dez pares (restart): ")
	next = makeEvenGenerator()
	for i := 0; i < vezes; i++ {
		fmt.Printf("\t%v", next())
	}
	fmt.Println("")

	fmt.Print("Dez primos: \t")
	prox := makePrimeGenerator()
	var prim uint
	for i := 1; i <= vezes; i++ {
		prim, _ = prox()
		fmt.Printf("\t%v", prim)
	}
	fmt.Println("")

	a, b := prox()
	fmt.Printf("%v \t%v \n", a, b)
	d, e := prox()
	fmt.Println(d)
	fmt.Printf("%T \t%T \t%T \n", b, d, e)

	fmt.Println("EXERCICIO EULER PROJECT:")
	inicio := time.Now()
	meuPrimo := makePrimeGenerator()
	var primoAtual uint
	var meusPrimos []uint
	for {
		primoAtual, _ = meuPrimo()
		if primoAtual >= 2000000 {
			_, meusPrimos = meuPrimo()
			break
		}
	}
	//fmt.Println(meusPrimos)
	fmt.Printf("%v \t %v \t %v \n", meusPrimos[len(meusPrimos)-3], meusPrimos[len(meusPrimos)-2], meusPrimos[len(meusPrimos)-1])
	meuTotal := addAllBelow(2000000, meusPrimos)
	fmt.Printf("%v\n", meuTotal)
	fim := time.Now()
	fmt.Println(fim.Sub(inicio))
}
