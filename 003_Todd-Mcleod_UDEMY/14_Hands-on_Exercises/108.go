package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range xi {
		fmt.Printf("%v \t %v \t %T \n", i, v, v)
	}
}

/*
	● Using a COMPOSITE LITERAL:
		● create a SLICE of TYPE int
		● assign these 10 VALUES
			42, 43, 44, 45, 46, 47, 48, 49, 50, 51
	● Range over the slice and print the values out.
		○ print out the VALUE and the TYPE

*/
