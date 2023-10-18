package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("Sending through the channel")
		for i := 0; i < 5; i++ {
			fmt.Println("Enviando i =", i+1)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			c <- i + 1
		}
		close(c) // o canal precisa ser fechado ou o programa retonará erro (all goroutines are asleep - deadlock)
	}()

	fmt.Println("Ranging over the channel")
	for v := range c {
		fmt.Println("Recebendo i =", v)
		fmt.Println(v)
	}

	fmt.Println("Final")
}
