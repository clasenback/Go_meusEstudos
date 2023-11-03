package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type pacote struct {
	thread   int64
	gerador  int
	worker   int
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
		for g := 0; g < geradores; g++ {
			geraPacotes(100, 20, g, c1)
			wg.Done()
		}
		close(c1)
	}()

	c2 := make(chan pacote)
	go func() {
		for w := 0; w < workers; w++ {
			worker(c1, c2, w)
			wg.Done()
		}
		close(c2)
	}()

	for pct := range c2 {
		fmt.Printf("g:%v w:%v %v\t%v\t%b\n", pct.gerador, pct.worker, pct.thread, pct.n, pct.fatorial)
	}
	wg.Wait()

	fmt.Println(time.Since(t))
}

func geraPacotes(n, max, g int, c chan<- pacote) {
	for i := 0; i < n; i++ {
		var pct pacote
		atomic.AddInt64(&contador, 1)
		pct.thread = contador
		pct.gerador = g
		pct.n = 1 + pct.thread%int64(max)
		c <- pct
	}
}

func worker(c1 <-chan pacote, c2 chan<- pacote, w int) {
	for pct := range c1 {
		pct.fatorial = fatorial(pct.n)
		pct.worker = w
		c2 <- pct
	}
}

func fanOutIn(c1, c2 chan pacote) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 pacote) {
			c2 <- doShit(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
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

func populate(c chan pacote) {
	var pct pacote
	for i := 0; i < 10; i++ {
		pct.nro = i
		pct.conteudo = i
		fmt.Printf("pacote %v criado: %v em %v\n", pct.nro, pct.conteudo, pct.tempo)
		c <- pct
	}
	close(c)
}
