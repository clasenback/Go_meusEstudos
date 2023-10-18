package main

import (
	"fmt"
	"time"
)

func whisper(l, r chan int) {
	l <- 1 + <-r
}

func main() {
	fmt.Println("caso não rode, 'go build'")

	t := time.Now()

	const n int = 1000000

	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 1; i <= n; i++ {
		right = make(chan int)
		//fmt.Println("Lançando whisper", i)
		go whisper(left, right)
		left = right
	}

	go func(c chan int) { c <- 0 }(right)
	//fmt.Println(<-right)
	fmt.Println(<-leftmost)
	fmt.Println(time.Since(t))

	t = time.Now()
	var soma int
	for i := 1; i <= n; i++ {
		soma++
	}
	fmt.Println(soma)
	fmt.Println(time.Since(t))
}

// 'go build' antes de rodar

/*
GOOGLE I/O 2012 - GO CONCURRENCY PATTERNS by Rob Pyke
https://www.youtube.com/watch?v=f6kdp27TYZs

ALGUMAS SUGESTÕES DE PROJETOS PARA ESTUDAR:
Chatroulette toy: 			tinyurl.com/gochatroulette
Load balancer: 				tinyurl.com/goloadbalancer
Concurrent prime sieve: 	tinyurl.com/gosieve
Concurrent power series: 	tinyurl.com/gopowerseries
*/
