package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type pacote struct {
	thread   int64
	n        int64
	fatorial int64
}

var contador int64

func main() {
	t := time.Now()

	geradores := 10
	workers := 6

	var wg sync.WaitGroup
	wg.Add(geradores + workers)

	c1 := make(chan pacote)
	go func() {
		for i := 0; i < geradores; i++ {
			geraPacotes(100, 20, c1)
			wg.Done()
		}
		close(c1)
	}()

	c2 := make(chan pacote)
	go func() {
		for w := 0; w < workers; w++ {
			worker(c1, c2)
			wg.Done()
		}
		close(c2)
	}()

	for pct := range c2 {
		fmt.Printf("%v\t%v\t%b\n", pct.thread, pct.n, pct.fatorial)
	}
	wg.Wait()

	fmt.Println(time.Since(t))
}

func geraPacotes(n, max int, c chan<- pacote) {
	for i := 0; i < n; i++ {
		var pct pacote
		atomic.AddInt64(&contador, 1)
		pct.thread = contador
		pct.n = 1 + pct.thread%int64(max)
		c <- pct
	}
}

func worker(c1 <-chan pacote, c2 chan<- pacote) {
	for pct := range c1 {
		pct.fatorial = fatorial(pct.n)
		c2 <- pct
	}
}

func fatorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	var total int64 = 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
