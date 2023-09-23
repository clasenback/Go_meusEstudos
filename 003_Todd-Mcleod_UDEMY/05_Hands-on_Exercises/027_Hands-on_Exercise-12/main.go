package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210
	fmt.Printf("%d \t %#b \t\t %#x \n", a, a, a)
	fmt.Printf("%d \t %#b \t\t %#x \n", b, b, b)
	fmt.Printf("%d \t %#b \t %#x \n", c, c, c)
}

/*
Hands-on exercise #12 (was #07)
- printf binary, decimal, & hexadecimal

write a program that uses print verbs to show the following numbers
● 747
● 911
● 90210
as
● decimal
● binary
● hexadecimal
*/
