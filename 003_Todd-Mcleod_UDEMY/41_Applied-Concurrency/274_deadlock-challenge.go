package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
