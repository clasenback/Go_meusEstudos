package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	go foo(c, wg)
	bar(c)

	fmt.Println("-----")
	passaCanal(c, c, wg)

	fmt.Println("Programa encerrado")
}

func foo(c chan<- int, wg sync.WaitGroup) {
	//wg.Add(1)
	fmt.Println("Enviando pelo canal")
	time.Sleep(time.Duration(rand.Intn(500)) * (time.Millisecond))
	c <- 42
	//wg.Done()
}

func bar(c <-chan int) {
	fmt.Println("Recebendo pelo canal")
	time.Sleep(time.Duration(rand.Intn(500)) * (time.Millisecond))
	fmt.Println(<-c)
}

func passaCanal(entra chan<- int, saida <-chan int, wg sync.WaitGroup) {
	go foo(entra, wg)
	bar(saida)
	//wg.Wait()
}
