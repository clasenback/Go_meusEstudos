/* import "fmt"

func main() {
	a := []int{0, 1, 2}
	testeste(a)
	fmt.Println(a)
}

func testeste(x []int) []int {
	return append(x, 1)
} */

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{11, 7, 5, 3, 2}
	for i, v := range a {
		fmt.Printf("i: %v \t v: %v \t a: %v %v \n", i, v, a[i], &a[i])
	}
	fmt.Println(sortSlice(a))
	for i, v := range a {
		fmt.Printf("i: %v \t v: %v \t a: %v %v \n", i, v, a[i], &a[i])
	}
	fmt.Println("------------")

	a = []int{11, 7, 5, 3, 2}
	for i, v := range a {
		fmt.Printf("i: %v \t v: %v \t a: %v %v \n", i, v, a[i], &a[i])
	}
	fmt.Println(copySlice(a))
	for i, v := range a {
		fmt.Printf("i: %v \t v: %v \t a: %v %v \n", i, v, a[i], &a[i])
	}
}

func sortSlice(xi []int) int {
	sort.Ints(xi)
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func copySlice(xi []int) int {
	x := make([]int, len(xi), cap(xi))
	copy(x, xi)
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

/*

 */
