package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("Programa encerrado")
}

func foo(c chan<- int) {
	fmt.Println("Enviando pelo canal")
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println("Recebendo pelo canal")
	fmt.Println(<-c)
}
