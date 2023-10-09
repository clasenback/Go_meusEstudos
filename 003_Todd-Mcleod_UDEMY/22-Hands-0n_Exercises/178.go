package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}

/*
In the provided code, there are variables that store VALUES of TYPE *int and TYPE *string. The
values stored are memory addresses. Using the variables provided, do the following:
	● print the VALUE stored in each variable
		○ these will be memory addresses
		● print the TYPE of each variable
		● print the data stored at memory locations
			○ dereference the stored memory address *
*/
