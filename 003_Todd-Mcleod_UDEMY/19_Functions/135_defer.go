/*
Este exercício verifica em que momento a função vinculada a defer é executada, se no momento em que é chamada pelo defer ou ao final do bloco em que o defer está inserido. O resultado indica que o defer é executado no momento em que é chamado e o seu valor guardado para ser apresentado ao final da execução do bloco.

A função makePrimeGenerator é utilizada apenas como forma de retardar a execução do código.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("Defer time:", mytimer())

	proximoPrimo := makePrimeGenerator()
	var primo uint
	for i := 0; i < 10000; i++ {
		primo, _ = proximoPrimo()
	}
	fmt.Println(primo)
	fmt.Println("Println time:", time.Now())

	for i := 0; i < 10000; i++ {
		primo, _ = proximoPrimo()
	}
	fmt.Println(primo)
	fmt.Println("Println time:", time.Now())
}

func mytimer() time.Time {
	return time.Now()
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
