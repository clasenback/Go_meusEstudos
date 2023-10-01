package main

import "fmt"

func main() {
	x := doShit(57, 48, mais)
	y := doShit(47, 9, menos)
	fmt.Printf("mais:   %v \t %T \n", mais, mais)
	fmt.Printf("menos:  %v \t %T \n", menos, menos)
	fmt.Printf("doShit: %v \t %T \n", doShit, doShit)
	fmt.Printf("x:      %v \t\t %T \n", x, x)
	fmt.Printf("y:      %v \t\t %T \n", y, y)
}

func mais(a, b int) int {
	return a + b
}

func menos(a, b int) int {
	return a - b
}

func doShit(a, b int, f func(int, int) int) int {
	return f(a, b)
}
