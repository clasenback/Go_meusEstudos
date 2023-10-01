package main

import "fmt"

func main() {
	pessoa := Pessoa{nome: "Alex"}
	agente := Agente{Pessoa: Pessoa{nome: "Monet"}, podeMatar: true}
	rainha := Rei{Pessoa: Pessoa{nome: "Lilica"}, governa: true}

	pessoa.falarNome()

	agente.falarNome()
	agente.atirar()

	rainha.falarNome()
	rainha.governar()

	Falar(pessoa)
	Falar(agente)
	Falar(rainha)
}

type Pessoa struct {
	nome string
}

func (p Pessoa) falarNome() {
	fmt.Println("Meu nome é", p.nome)
}

/*
func (p *Pessoa) tornarAgente() {
	p = Agente{Pessoa: p, podeMatar: true}
} */

type Agente struct {
	Pessoa
	podeMatar bool
}

func (a Agente) falarNome() {
	a.Pessoa.falarNome()
	fmt.Println("e eu sou um agente secreto!")
}
func (a Agente) atirar() {
	if a.podeMatar {
		fmt.Println(a.nome, "atirou em você.")
	}
}

type Rei struct {
	Pessoa
	governa bool
}

func (r Rei) falarNome() {
	fmt.Println("Eu sou o Rei", r.nome)
}

func (r Rei) governar() {
	fmt.Println("Faça o que eu digo, não faça o que eu faço.")
}

type Humano interface {
	falarNome()
}

func Falar(h Humano) {
	fmt.Println("========")
	h.falarNome()
	fmt.Println("========")
}
