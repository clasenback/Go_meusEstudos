package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	x, y, raio float64
}

type Retangulo struct {
	x1, y1, x2, y2 float64
}

func (c *Circulo) Area() float64 {
	return c.raio * c.raio * math.Pi
}

func (c *Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (r *Retangulo) Altura() float64 {
	return Distancia(r.x1, r.y1, r.x1, r.y2)
}

func (r *Retangulo) Base() float64 {
	return Distancia(r.x1, r.y1, r.x2, r.y1)
}

func (r *Retangulo) Area() float64 {
	return r.Base() * r.Altura()
}

func (r *Retangulo) Perimetro() float64 {
	return 2*r.Base() + 2*r.Altura()
}

func Distancia(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {
	c := Circulo{10, 5, 8}
	fmt.Println("Círculo")
	fmt.Printf("Área: \t %.2f \t Perímetro: \t %.2f \n", c.Area(), c.Perimetro())
	r := Retangulo{6, 8, 5, 10}
	fmt.Println("Retangulo")
	fmt.Printf("Área: \t %.2f \t\t Perímetro: \t %.2f \n", r.Area(), r.Perimetro())
}
