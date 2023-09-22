package main

import (
	"fmt"
	"math/rand"
)

func doSomeShitWithIf(i int) {
	var x int = rand.Intn(10)
	var y int = rand.Intn(10)
	fmt.Printf("%v \t O valor de x é %v e de y é %v. ", i, x, y)
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

func main() {
	for i := 1; i <= 100; i++ {
		doSomeShitWithIf(i)
	}
}

/*
● there are two parts ot this hands on exercise
	○ create a program that has a loop that prints every number from 0 to 99
	○ modify the program from the previous hands on exercise to run 100 times
*/
