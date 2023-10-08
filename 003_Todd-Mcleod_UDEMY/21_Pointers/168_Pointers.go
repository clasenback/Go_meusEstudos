package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*&x)
	fmt.Printf("%T\t%v\n", &x, &x)

	px := &x
	*px = 43
	fmt.Println(x, px)
}
