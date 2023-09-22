package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println(i, xi[i], v)
	}
}

/*
xi := []int{42, 43, 44, 45, 46, 47}
● use a for range loop to print each value and the index position of each value
*/
