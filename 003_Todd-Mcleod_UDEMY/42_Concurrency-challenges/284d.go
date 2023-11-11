package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type pacote struct {
	pacote   int
	random   int
	fatorial int
}

func main() {
	t := time.Now()

	c1 := make(chan pacote)
	c2 := make(chan pacote)

	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			wg1.Add(1)
			go func(pkg int) {
				var pct pacote
				pct.pacote = pkg
				pct.random = rand.Intn(21)
				c1 <- pct
				wg1.Done()
			}(i)
		}
		wg1.Done()
		wg1.Wait()
		close(c1)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		for pct := range c1 {
			pct.fatorial = fatorial(pct.random)
			c2 <- pct
		}
		wg2.Done()
		wg2.Wait()
		close(c2)
	}()

	for pct := range c2 {
		fmt.Printf("Pacote: %v\tConteúdo: %v\tFatorial: %v\n", pct.pacote, pct.random, pct.fatorial)
	}

	fmt.Println("Execução em: ", time.Since(t))
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}
