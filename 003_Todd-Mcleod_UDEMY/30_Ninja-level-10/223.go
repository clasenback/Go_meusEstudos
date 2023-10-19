package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	f := func(c chan int, i int) {
		c <- i
	}

	go f(c, 42)
	fmt.Println(<-c)

	cc := make(chan int, 2)
	go f(cc, 43)
	go f(cc, 44)
	fmt.Println(<-cc)
	fmt.Println(<-cc)

	fmt.Printf("c é do tipo %T\n", c)
	fmt.Printf("cc é do tipo %T", cc)
}
