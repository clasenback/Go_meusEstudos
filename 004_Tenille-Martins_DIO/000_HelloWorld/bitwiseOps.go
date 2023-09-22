package main

import "fmt"

func main() {
	fmt.Printf("%d \t %#b \n", 1, 1)
	fmt.Printf("%d \t %#b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %#b \n", 1<<5-1, 1<<5-1)
}
