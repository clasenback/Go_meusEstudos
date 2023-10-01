package main

import "fmt"

func main() {
	fmt.Println(fatorial(6))
}

func fatorial(n int) int {
	if n == 1 {
		fmt.Println("1 =")
		return 1
	} else {
		fmt.Print(n, " * ")
		return n * fatorial(n-1)
	}
}
