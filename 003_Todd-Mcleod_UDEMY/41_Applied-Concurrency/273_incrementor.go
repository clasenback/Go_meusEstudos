package main

import "fmt"

func main() {
	c1 := incrementor("A 😎")
	c2 := incrementor("😁 B")
	c11 := puller(c1)
	c21 := puller(c2)
	fmt.Println("soma:", <-c11+<-c21)
}

//gerador
func incrementor(s string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(s, i)
			out <- 1
		}
		close(out)
	}()
	return out
}

//worker
func puller(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var soma int
		for n := range in {
			soma += n
			fmt.Println(n, soma)
		}
		out <- soma
		close(out)
	}()
	return out
}
