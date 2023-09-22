package main

import "fmt"

type retangulo struct {
	altura, largura float64
}

func (r retangulo) perimetro() float64 {
	return r.altura * r.largura
}

func (r retangulo) area() float64 {
	return 2.0 * (r.altura + r.largura)
}

func main() {
	meuRet := retangulo{3, 4}
	fmt.Println(meuRet.perimetro())
	fmt.Println(meuRet.area())
}
