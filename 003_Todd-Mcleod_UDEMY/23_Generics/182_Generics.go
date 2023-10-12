package main

import "fmt"

func main() {
	var a int = 3
	var b float32 = 3.14
	fmt.Println(add1(a, a))
	fmt.Println(add1(b, b))
	// fmt.Println(add(a, b))

	fmt.Println(add2(b, b))

	var c meuInteiro = 4
	s := add3(c, c)
	fmt.Printf("%v\t%T\n", s, s)

}

func add1[T float32 | int](a, b T) T {
	return a + b
}

type meuTipo interface {
	float32 | int
}

func add2[T meuTipo](a, b T) T {
	return a + b
}

type meuTipo2 interface {
	~float32 | ~int
}

//add3 recebe a e b de quailquer tipo rastreável como float32 ou int e retorna a soma com o mesmo tipo que recebeu. Ambos parametros precisam ser do mesmo tipo. 
func add3[T meuTipo2](a, b T) T {
	return a + b
}

type meuInteiro int
