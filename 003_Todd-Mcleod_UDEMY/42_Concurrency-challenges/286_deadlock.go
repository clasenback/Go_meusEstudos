package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	//FAN OUT
	//Multiple functions reading from the same chanel until that chanel is closed
	//Distribute work across multiple functions (10 goroutines) that all read from in.
	xc := fanOut(in, 10)
	fmt.Println("Tamanho de xc:", len(xc))
	//FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	for pacote := range merge(xc...) {
		fmt.Printf("i: %v\tj: %v\tfan: %v\tfatorial: %v\n", pacote.i, pacote.j, pacote.fan, pacote.fatorial)
	}
}

type dados struct {
	i        int
	j        int
	fatorial int
	fan      int
}

// OK, sem bug
func gen() <-chan dados {
	out := make(chan dados)
	go func() {
		for i := 0; i < 3; i++ {
			for j := 3; j < 8; j++ {
				var pct dados
				pct.i = i
				pct.j = j
				out <- pct
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan dados, fan int) <-chan dados {
	out := make(chan dados)
	go func() {
		for pkg := range in {
			pkg.fatorial = fact(pkg.j)
			pkg.fan = fan
			out <- pkg
		}
		close(out) //erro aqui
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func fanOut(in <-chan dados, n int) []<-chan dados {
	xc := make([]<-chan dados, 0)
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in, i))
	}
	return xc
}

func merge(cs ...<-chan dados) <-chan dados {
	var wg sync.WaitGroup
	out := make(chan dados)

	// função output
	output := func(c <-chan dados) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all output goroutines are
	// done. This must start after wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- This code throws an error: fatal error: all goroutines are asleep - deadlock!
-- fix this code!
*/
