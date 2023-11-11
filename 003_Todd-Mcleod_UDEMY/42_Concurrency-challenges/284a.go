package main

import (
	"fmt"
	"math/rand"
)

type pacote struct {
	random int
	c1     int
	c2     int
}

func main() {
	c1 := make(chan pacote)
	c2 := make(chan pacote)

	var pct pacote
	pct.random = rand.Intn(20) + 1
	fmt.Printf("Início:\t\t%v\t%v\t%v\n", pct.random, pct.c1, pct.c2)
	go func() {
		pct.c1 = rand.Intn(20) + 1
		fmt.Printf("1º goroutine:\t%v\t%v\t%v\n", pct.random, pct.c1, pct.c2)
		c1 <- pct
	}()

	go func() {
		pct := <-c1
		pct.c2 = rand.Intn(20) + 1
		fmt.Printf("2º goroutine:\t%v\t%v\t%v\n", pct.random, pct.c1, pct.c2)
		c2 <- pct
	}()

	fmt.Println(<-c2)
}
