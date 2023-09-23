package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x int = rand.Intn(250)
	fmt.Println("x:", x)
	switch {
	case x <= 100:
		fmt.Println("entre 0 e 100")
	case x <= 200:
		fmt.Println("entre 101 e 200")
	default:
		fmt.Println("entre 201 e 250")
	}
}

/*
Modify the previous program to use a switch statement
*/
