package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	//tipo string = "Quadrado"
	lado float64
}

type circulo struct {
	//tipo string = "Círculo"
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Println(g)
	fmt.Printf("Área: \t\t %v \n", g.area())
	fmt.Printf("Perimetro: \t %v \n", g.perimetro())
}

func main() {
	meuQ := quadrado{lado: 10}
	meuC := circulo{raio: 5}
	medir(meuQ)
	medir(meuC)
}
