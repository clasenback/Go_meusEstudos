package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	fim := make(chan int)

	go send(par, impar, fim)

	receive(par, impar, fim)

	fmt.Println("Finalizado")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Canal par:", v)
		case v := <-o:
			fmt.Println("Canal impar:", v)
		case v := <-q:
			fmt.Println("Canal fim:", v)
			return
		}
	}
}
func send(e, o, q chan<- int) {
	for i := 0; i < 13; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
