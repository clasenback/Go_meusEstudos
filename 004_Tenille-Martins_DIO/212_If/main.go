package main

import "fmt"

func main() {
	for i := 0; i < 25; i++ {
		if i%2 > 0 {
			fmt.Printf("nro %v é impar. \n", i)
		} else {
			fmt.Printf("nro %v é par. \n", i)
		}
	}
	fmt.Println()
}
