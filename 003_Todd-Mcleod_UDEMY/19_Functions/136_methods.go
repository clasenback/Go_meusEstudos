package main

import (
	"fmt"
	"math"
)

func main() {
	meuC := Circulo{10}
	fmt.Printf("%.2f \t%.2f \t%.2f \n", meuC.raio, meuC.Area(), meuC.Perimetro())
	meuC.Escala(5)
	fmt.Printf("%.2f \t%.2f \t%.2f \n", meuC.raio, meuC.Area(), meuC.Perimetro())
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (c *Circulo) Escala(escala float64) {
	c.raio *= escala
}
