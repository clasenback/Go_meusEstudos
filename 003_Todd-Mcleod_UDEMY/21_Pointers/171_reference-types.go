package main

import "fmt"

func main() {
	x := 42
	plus(&x)
	fmt.Println(x)

	xi := []int{11, 22, 33, 44}
	altera(xi)
	fmt.Println(xi)

	mp := make(map[string]int)
	mp["Alex"] = 44
	alteramp(mp)
	fmt.Println(mp)
}

func plus(p *int) {
	*p++
}

func altera(xi []int) {
	xi[0] = 99
}

func alteramp(mp map[string]int) {
	mp["Alex"] = 1979
	mp["Norma"] = 1981
}
