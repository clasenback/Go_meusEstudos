package main

import "fmt"

func fatorial(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		fmt.Print("2 * 1 = \n")
		return 2
	default:
		fmt.Print(n, " * ")
		return n * fatorial(n-1)
	}
}

func main() {
	var x int = 4

	cl := func() int {
		x++
		return x
	}

	fmt.Println("O fatorial de", x, "+ 1 é")
	fmt.Println(fatorial(cl()))
}
