package main

import (
	"fmt"
	"math/rand"
)

func doSomeShitWithSwitch(i int) {
	var x int = rand.Intn(5)
	fmt.Printf("%v \t", i)
	switch x {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Três")
	case 4:
		fmt.Println("Quatro")
	default:
		fmt.Println("Cinco")
	}
}

func main() {
	for i := 0; i < 42; i++ {
		doSomeShitWithSwitch(i)
	}
}

/*
● Create one random int between 0 inclusive and 5 exclusive
	○ assign the value to a variable with the identifier x
● Use a switch statement to print a description of the variable and value
● run the code 42 times and print the iteration number
*/
