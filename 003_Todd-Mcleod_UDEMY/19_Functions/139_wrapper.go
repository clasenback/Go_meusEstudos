package main

import (
	"fmt"
	"time"
)

func main() {
	Cronometrar(Dormir)
}

func Cronometrar(fn func() int) {
	// wrapper function
	t := time.Now()
	res := fn()
	fmt.Printf("A função foi executada em %v \t %v \n", time.Since(t), res)
}

func Dormir() int {
	time.Sleep(2 * time.Second)
	fmt.Println("Função DORMIR executada.")
	return 1
}
