package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x int = rand.Intn(10)
	var y int = rand.Intn(10)
	fmt.Printf("O valor de x é %v e de y é %v \n", x, y)
	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}
}

/*
● Create 2 random ints between 0 inclusive and 10 exclusive
	○ assign them to variables with the identifiers x and y
● Print their values
● use an if statement to print the correct description
	○ x and y are both less than 4
	○ x and y are both greater than 6
	○ x is greater than or equal to 4 and less than or equal to 6
	○ y is not 5
	○ none of the previous cases were met
*/
