package main

import (
	"fmt"
	"time"
)

func main() {
	Cronometrar(Dormir)
}

func Cronometrar(fn func()) {
	// wrapper function
	t := time.Now()
	fn()
	fmt.Println("A função foi executada em", time.Since(t))
}

func Dormir() {
	time.Sleep(2 * time.Second)
	fmt.Println("Função DORMIR executada.")
}
