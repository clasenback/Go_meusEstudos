package main

import "fmt"

func main() {
	f()
}

var f = func() {
	fmt.Println("essa é a variável f")
}
