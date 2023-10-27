package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 7 =", mySum(4, 7))
	fmt.Println("5 + 9 =", mySum(5, 9))
	fmt.Println("2 - 3 =", mySub(2, 3))
	fmt.Println("4 - 7 =", mySub(4, 7))
	fmt.Println("5 - 9 =", mySub(5, 9))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func mySub(xi ...int) int {
	sub := 0
	for _, v := range xi {
		sub -= v
	}
	return sub
}
