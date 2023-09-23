package main

import "fmt"

func main() {
	var a int8 = 127
	var b uint8 = 255
	fmt.Printf("%d \t %#b \t %T \n", a, a, a)
	fmt.Printf("%d \t %#b \t %T \n", b, b, b)
}

/*
Hands-on exercise #13 (was #08)
- signed and unsigned int

write a program that declares two variables
● one variable to store a VALUE of TYPE int8
	○ assign to it the largest number possible, then print it
● one variable to store a VALUE of TYPE uint8
	○ assign to it the largest number possible, then print it

	https://go.dev/ref/spec#Numeric_types
	uint8       the set of all unsigned  8-bit integers (0 to 255)
	int8        the set of all signed  8-bit integers (-128 to 127)

*/
