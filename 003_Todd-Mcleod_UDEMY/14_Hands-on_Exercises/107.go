package main

import "fmt"

func main() {
	var ax [5]int
	for i, _ := range ax {
		ax[i] = i * 3
	}
	for i, v := range ax {
		fmt.Printf("%v \t %T \t %v \t %T \n", i, i, v, v)
	}
}

/*
	● Using a COMPOSITE LITERAL:
		● create an ARRAY which holds 5 VALUES of TYPE int
		● assign VALUES to each index position.
	● Range over the array and print the values out.
		○ print out the VALUE and the TYPE
*/
