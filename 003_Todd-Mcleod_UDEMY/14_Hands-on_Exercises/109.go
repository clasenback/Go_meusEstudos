package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xxi := [][]int{xi[:5], xi[5:], xi[2:7], xi[1:6]}
	for i, v := range xxi {
		fmt.Printf("%v \t %v \t %T \n", i, v, v)
	}
}

/*
	Using the code from the previous example,
	use SLICING to create the following new slices
	which are then printed:
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
*/
