package main

import (
	"fmt"
)

func main() {
	quit := make(chan bool)
	data := gen(quit)
	receive(data, quit)

	fmt.Println("about to exit")
}

func gen(q chan<- bool) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- true
		close(c)
	}()

	return c
}

func receive(c <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case v := <-q:
			fmt.Println("quit:", v)
			return
		}
	}
}
