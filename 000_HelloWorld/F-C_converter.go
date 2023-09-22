package main

import "fmt"

func main() {
	fmt.Println("Digite quantos pés: ")
	var pes float64
	fmt.Scanf("%f", &pes)
	const metros float64 = 0.3048
	var metro float64
	metro = pes * metros
	fmt.Println(metro)

	var i int8
	i = 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}

	i = 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println(i)
		}
		i++
	}

	i = 1
	for i < 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "- 3 e 5")
		} else if i%3 == 0 {
			fmt.Println(i, "- 3")
		} else if i%5 == 0 {
			fmt.Println(i, "- 5")
		}
		i++
	}
}
