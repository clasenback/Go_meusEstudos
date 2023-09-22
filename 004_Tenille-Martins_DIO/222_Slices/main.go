package main

import "fmt"

func main() {
	arr := [7]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(arr)
	fatia := make([]int, 4)
	fmt.Println(fatia)
	fatia = arr[2:6]
	fmt.Println(fatia)

	fatia1 := []int{2, 4, 6}
	fatia2 := append(fatia1, 8, 10)
	fmt.Println(fatia1, fatia2)
	//o tamanho do slice pode aumentar indefinidamente
	fatia2 = append(fatia2, 12, 14)
	fmt.Println(fatia2)

	fatia3 := make([]int, 1)
	copy(fatia3, fatia1)
	fmt.Println(fatia1, fatia3)
}
