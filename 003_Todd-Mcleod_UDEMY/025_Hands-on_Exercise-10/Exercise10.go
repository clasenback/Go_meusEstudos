package main

import "fmt"

func main() {
	fmt.Println("var for zero value")
	var a int
	fmt.Println(a)

	fmt.Println("short declaration operator")
	b := 42.0
	fmt.Println(b)

	fmt.Println("multiple initializations")
	c, d := "Alex", true
	fmt.Println(c, d)

	fmt.Println("var when specificity is required")
	var e float64 = 3.1416
	fmt.Printf("%v \t %T \n", e, e)

	fmt.Println("blank identifier")
	f, _, g := "a", "b", "c"
	fmt.Println(f, g)
}

/*
Hands-on exercise #10 (was #05)
- zero value, :=, type specificity, blank identifier
write a program that uses the following:
● var for zero value
● short declaration operator
● multiple initializations
● var when specificity is required
● blank identifier
print items as necessary to make the program interesting
*/
