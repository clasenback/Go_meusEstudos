package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	fim := make(chan bool)

	go send(par, impar, fim)

	receive(par, impar, fim)

	fmt.Println("Finalizado")
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("Canal par:", v)
		case v := <-o:
			fmt.Println("Canal impar:", v)
		case v, ok := <-q: // curioso: a variável v recebe int e bool do canal
			fmt.Println("Canal fim:", v, "\tok:", ok)
			return
		}
	}
}
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 13; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- true
}
