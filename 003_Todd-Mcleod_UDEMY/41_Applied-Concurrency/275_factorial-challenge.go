package main

import "fmt"

func main() {
	fac := factorial(6)
	fmt.Println(<-fac)
}

func factorial(n int) chan int {
	ans := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		ans <- total
		close(ans)
	}()
	return ans
}
