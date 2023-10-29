package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	d := make(chan bool)
	grs := []string{"A", "B", "C", "D"}

	for _, gr := range grs {
		go doShit(gr, c, d)
	}

	go func() {
		for n := range grs {
			fmt.Println("Goroutine", n, "concluída:", <-d)
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Encerrando o programa")
}

func doShit(s string, c chan<- string, d chan<- bool) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		c <- fmt.Sprintf("%v: %v", s, i)
	}
	d <- true
}
