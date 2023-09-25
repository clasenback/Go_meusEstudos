package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11}
	b := a
	c := make([]int, 5)
	copy(c, a)
	fmt.Println(len(a), len(b), len(c))
	for i, v := range a {
		fmt.Printf("i: %v \t v: %v \t a: %v %v \t b: %v %v \t c: %v %v\n", i, v, a[i], &a[i], b[i], &b[i], c[i], &c[i])
	}
	fmt.Println("------------")
	a = append(a, 13)
	b = append(b, 15)
	c = append(c, 17)
	fmt.Println(len(a), len(b), len(c))
	for i, v := range a {
		fmt.Printf("i: %v \t v: %v \t a: %v %v \t b: %v %v \t c: %v %v\n", i, v, a[i], &a[i], b[i], &b[i], c[i], &c[i])
	}
}

/*
	os slices A e B apontam para a mesma posição na memória
	o slice C aponta para uma posição nova na memória, para onde é copiado o conteúdo de A
	após append(A), os ponteiros &A[] são alocados para uma nova posição de memória,
	A e B não apontam mais para a mesma posição na memória, portanto
*/
