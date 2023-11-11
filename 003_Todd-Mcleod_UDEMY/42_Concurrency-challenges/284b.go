package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type pacote struct {
	random int
	c1     int
	c2     int
}

func main() {
	c1 := make(chan pacote)
	c2 := make(chan pacote)

	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		for i := 0; i < 3; i++ {
			var pct pacote
			pct.random = rand.Intn(20) + 1
			pct.c1 = i + 1
			fmt.Printf("1º goroutine:\t%v\t%v\t%v\n", pct.random, pct.c1, pct.c2)
			c1 <- pct
		}
		wg1.Done()
		wg1.Wait()
		close(c1)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		for pct := range c1 {
			pct.c2 = pct.c1
			fmt.Printf("2º goroutine:\t%v\t%v\t%v\n", pct.random, pct.c1, pct.c2)
			c2 <- pct
		}
		wg2.Done()
		wg2.Wait()
		close(c2)
	}()

	for pct := range c2 {
		fmt.Println(pct)
	}
}
