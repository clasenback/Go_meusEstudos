package main

import "fmt"

func main() {
	i := 0
	for i < 13 {
		if i%2 == 0 && i%3 == 0 {
			fmt.Println(i)
		}
		i++
	}
}
