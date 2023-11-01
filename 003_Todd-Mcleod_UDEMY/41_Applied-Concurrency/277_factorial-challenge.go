package main

import (
	"fmt"
	"time"
)

func main() {
	nros := gen(100)
	fac := fact(nros)
	for v := range fac {
		fmt.Printf("%v\t%b\t%v\n", v.n, v.fac, v.t)
	}
}

func gen(n uint64) chan uint64 {
	out := make(chan uint64)
	go func() {
		var i uint64
		for i = 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func fact(in chan uint64) chan facto {
	out := make(chan facto)
	go func() {
		for n := range in {
			t := time.Now()
			var res facto
			res.fac = 1
			res.n = n
			for i := n; i > 0; i-- {
				res.fac *= i
			}
			res.t = time.Since(t)
			out <- res
		}
		close(out)
	}()
	return out
}

type facto struct {
	n   uint64
	fac uint64
	t   time.Duration
}
