package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println("soma:", n)
	}
}

//gerador
func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			out <- i
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
