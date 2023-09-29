/*INTRODUCING GO: EXERCÍCIOS DO CAPÍTULO 6*/

package main

import "fmt"

func addAll(xi []int) (total int) {
	/* 	sum is a function that takes a slice of numbers
	and adds them together. What would its function
	signature look like in Go?
	*/
	for _, v := range xi {
		total += v
	}
	return
}

func half(n int) (int, bool) {
	/* Write a function that takes an integer and halves it
	and returns true if it was even or false if it was odd.
	For example, half(1) should return (0, false) and half(2)
	should return (1, true).
	*/
	return n / 2, n%2 == 0
}

func greatest(args ...int) (n int) {
	/* Write a function with one variadic parameter that finds
	the greatest number in a list of numbers. */
	n = args[0]
	for _, v := range args {
		if v > n {
			n = v
		}
	}
	return
}

func makeOddGenerator() func() uint {
	n := uint(1)
	return func() (m uint) {
		m = n
		n += 2
		return
	}

}

func testePointers() func() *int {
	vf := 1
	vf++
	return func() *int {
		pf := &vf
		fmt.Printf("a variável vf é %v e está alocada no endereço %v", vf, pf)
		return pf
	}
}

func fibo(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func square(x *float64) {
	*x = *x * *x
}

func one(xPtr *int) {
	*xPtr++
}

func swap(aaPtr, bbPtr *int) {
	*aaPtr, *bbPtr = *bbPtr, *aaPtr
}

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	//EXERCÍCIO 1
	myslice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(addAll(myslice))

	//EXERCÍCIO 2
	fmt.Println(half(6))

	//EXERCÍCIO 3
	fmt.Println(greatest(2, 4, 8, 15, 2, 4, 17, 9))

	//EXERCÍCIO 4
	proxImpar := makeOddGenerator()
	fmt.Print(proxImpar(), "\t")
	fmt.Print(proxImpar(), "\t")
	fmt.Print(proxImpar(), "\n")

	//EXERCÍCIO 5
	for i := 1; i < 12; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Print("\n")
	for i := 1; i < 12; i++ {
		fmt.Print(fibo(uint(i)), "\t")
	}
	fmt.Print("\n")

	//EXERCÍCIO 7
	teste := 10
	fmt.Printf("A variável teste tem valor %v, seu tipo é %T, seu endereço na memória aponta para %v e o tipo do seu ponteiro é %T \n", teste, teste, &teste, &teste)

	//EXERCÍCIO 10
	x := 1.5
	square(&x)
	// a função square é aplicada ao que estiver contido no endereço da variável x
	// não é o conteúdo de x que vai para a função, é a função que é aplicada ao conteúdo.
	fmt.Println(x)

	//PONTEIROS
	fmt.Println()
	i := 42
	fmt.Println("VARIÁVEL i")
	fmt.Println("a variável i foi criada")
	fmt.Printf("a variável i foi alocada em %v e esse endereço recebe valores do tipo %T \n", &i, &i)
	fmt.Printf("a variável i armazena o valor %v de tipo %T \n", i, i)

	p := &i
	fmt.Println("VARIÁVEL p - PONTEIRO")
	fmt.Printf("o endereço da variável i foi armazenado no ponteiro p \n")
	fmt.Printf("o ponteiro p foi alocado em %v e esse endereço recebe valores do tipo %T \n", &p, &p)
	fmt.Printf("o endereço apontado por p é %v e esse endereço recebe valores do tipo %T \n", p, p)
	fmt.Printf("o conteúdo no endereço apontado por p é %v, de tipo %T \n", *p, *p)

	q := &p
	fmt.Println("VARIÁVEL q - PONTEIRO PARA PONTEIRO")
	fmt.Printf("o endereço do ponteiro p foi armazenado no ponteiro q \n")
	fmt.Printf("o ponteiro q foi alocado em %v e esse endereço recebe valores do tipo %T \n", &q, &q)
	fmt.Printf("o endereço apontado por q é %v e esse endereço recebe valores do tipo %T \n", q, q)
	fmt.Printf("o conteúdo no endereço apontado por q é %v, de tipo %T \n", *q, *q)
	fmt.Printf("o conteúdo no endereço apontado por q e p é %v, de tipo %T \n", **q, **q)

	fmt.Println("VARIÁVEL ff - PONTEIRO PARA CLOSURE")
	fPtr := testePointers()
	ff := fPtr()
	fmt.Println()
	fmt.Println(ff, *ff, fPtr)
	*ff = 42
	fmt.Println(ff, *ff, fPtr)
	ff = fPtr()
	fmt.Println()
	fmt.Println(ff, *ff, fPtr)

	//EXERCÍCIO 11
	/* 	Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)
	should give you x=2 and y=1). */
	aa, bb := 1, 2
	fmt.Println(aa, bb)
	swap(&aa, &bb)
	fmt.Println(aa, bb)

	//panic("PANIC")
}
