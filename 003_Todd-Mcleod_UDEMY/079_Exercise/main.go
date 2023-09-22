package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x int = rand.Intn(10)
	var y int = rand.Intn(10)
	fmt.Printf("O valor de x é %v e de y é %v \n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}
}
