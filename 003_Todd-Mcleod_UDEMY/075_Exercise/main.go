package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x int = rand.Intn(250)
	fmt.Println("x:", x)
	if x <= 100 {
		fmt.Println("entre 0 e 100")
	} else if x <= 200 {
		fmt.Println("entre 101 e 200")
	} else {
		fmt.Println("entre 201 e 250")
	}
}

/*
create a program that uses both SEQUENTIAL and CONDITIONAL control flow.
Your program should do the following
	○ create a random int between 0 andrand
	○ store the value of that int in a variable with the identifier of x
		■ func Intn(n int) int
			● rand.Intn()
	○ print out the name and value of the variable
	○ use an if statement to do the following
		■ if the value is between 0 and 100
			● print between 0 and 100
		■ if the value is between 101 and 200
			● print between 101 and 200
		■ if the value is between 201 and 250
			● print between 201 and 250
			● re: inclusive, exclusive – does rand.Intn()
	○ include zero in its output?
	○ include 250 in its output?
	○ show this in code using the numbers 0 - 3
*/
