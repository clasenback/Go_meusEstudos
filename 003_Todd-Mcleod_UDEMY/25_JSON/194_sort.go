package main

import (
	"fmt"
	"sort"
)

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("como criado\t", people)
	sort.Sort(ByAge(people))
	fmt.Println("por idade\t", people)
	sort.Sort(ByPerson(people))
	fmt.Println("por nome\t", people)
}

type Person struct {
	First string
	Age   int
}

// tipo ByAge (e seus métodos) implementa o tipo sort.Interface e permite que a função sort.Sort seja utilizada com a estrutura de dados do exercício.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// tipo ByPerson (e seus métodos) implementa o tipo sort.Interface e permite que a função sort.Sort seja utilizada com a estrutura de dados do exercício. O método Less é a única alteração, comparativamente ao tipo ByAge.
type ByPerson []Person

func (r ByPerson) Len() int           { return len(r) }
func (r ByPerson) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByPerson) Less(i, j int) bool { return r[i].First < r[j].First }
