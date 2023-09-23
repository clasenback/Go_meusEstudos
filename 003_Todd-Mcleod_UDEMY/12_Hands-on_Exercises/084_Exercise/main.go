package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

/*
● use modulus and the continue statement in a loop to print all ODD numbers
● joke about the programmer and infinite loops
*/
