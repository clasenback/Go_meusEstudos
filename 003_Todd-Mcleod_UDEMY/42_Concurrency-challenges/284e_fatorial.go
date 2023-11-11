package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type pacote struct {
	pacote   int
	worker   int
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
		for i := 0; i < 1000; i++ {
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
		for w := 0; w < 8; w++ {
			wg2.Add(1)
			go func(worker int) {
				for pct := range c1 {
					pct.worker = worker
					//espera(pct.random)
					pct.fatorial = fatorial(pct.random)
					c2 <- pct
				}
				wg2.Done()
			}(w)
		}
		wg2.Done()
		wg2.Wait()
		close(c2)
	}()

	for pct := range c2 {
		fmt.Printf("Pacote: %v\tConteúdo: %v\tWorker: %v\tFatorial: %b\n", pct.pacote, pct.random, pct.worker, pct.fatorial)
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

func espera(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}
