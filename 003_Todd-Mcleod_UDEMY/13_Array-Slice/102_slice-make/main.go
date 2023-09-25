package main

import "fmt"

func main() {
	si := make([]int, 5, 5)
	fmt.Println(si, len(si), cap(si), &si[0])
	si = append(si, 1)
	fmt.Println(si, len(si), cap(si), &si[0])
	si = append(si, 1)
	fmt.Println(si, len(si), cap(si), &si[0])
	si = append(si, 2, 3, 4, 5)
	fmt.Println(si, len(si), cap(si), &si[0])
	si = append(si, si...)
	fmt.Println(si, len(si), cap(si), &si[0])
	x := 5
	fmt.Println(x, &x)
}
