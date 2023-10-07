package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)
		fmt.Println("goroutine", n, "loop", i, "em", time.Millisecond*amt)
	}
	fmt.Println("Goroutine", n, "concluida em", time.Since(t0))
}
