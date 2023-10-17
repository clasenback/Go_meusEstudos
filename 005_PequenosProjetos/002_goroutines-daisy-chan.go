package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	const n = 5
	var wg sync.WaitGroup
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	wg.Add(n)
	for i := 0; i < n; i++ {
		right = make(chan int)
		go func(l, r chan int) {
			left <- 1 + <-right
			fmt.Println()
			wg.Done()
		}(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	final := <-leftmost
	fmt.Println(final)
	wg.Wait()
	fmt.Println(time.Since(t))
}

/*
GOOGLE I/O 2012 - GO CONCURRENCY PATTERNS by Rob Pyke
https://www.youtube.com/watch?v=f6kdp27TYZs

ALGUMAS SUGESTÕES DE PROJETOS PARA ESTUDAR:
Chatroulette toy: 			tinyurl.com/gochatroulette
Load balancer: 				tinyurl.com/goloadbalancer
Concurrent prime sieve: 	tinyurl.com/gosieve
Concurrent power series: 	tinyurl.com/gopowerseries
*/
