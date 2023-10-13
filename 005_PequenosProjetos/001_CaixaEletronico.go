package main

import (
	"fmt"
	"math/rand"
)

func main() {
	meuCaixa := iniciarCaixaAleatorio(100)
	caixaInicial := meuCaixa
	fmt.Printf("Total em caixa: $%v\n", somaNotas(meuCaixa))

	var meuSaque saque
	var saquesOK, saquesParciais, caixaInsuficiente, caixaInsufSeguidos int
	var extratos []extrato

	for {
		queroSacar := valor(rand.Intn(1000))
		fmt.Println("")
		fmt.Printf("VALOR DO SAQUE: $%v\n", queroSacar)

		fmt.Println("")
		fmt.Println("==== INICIAR CONTAGEM DE NOTAS ====")
		meuSaque, meuCaixa = requisitarNotas(queroSacar, meuCaixa)

		fmt.Println("")
		fmt.Println("==== CAIXA RESTANTE ====")
		contarNotas(meuCaixa)
		fmt.Printf("Total em caixa: $%v\n", somaNotas(meuCaixa))

		fmt.Println("")
		fmt.Println("==== ENTREGAR NOTAS ====")
		contarNotas(meuSaque)
		fmt.Printf("Saque requisitado:\t$%v\nSaque realizado:\t$%v\n", queroSacar, somaNotas(meuSaque))

		switch {
		case queroSacar == somaNotas(meuSaque):
			saquesOK++
			caixaInsufSeguidos = 0
			extratos = append(extratos, extrato{requisitado: queroSacar, sacado: meuSaque, flag: 1, emCaixa: somaNotas(meuCaixa)})
		case queroSacar < somaNotas(meuCaixa):
			saquesParciais++
			caixaInsufSeguidos = 0
			extratos = append(extratos, extrato{requisitado: queroSacar, sacado: meuSaque, flag: 0, emCaixa: somaNotas(meuCaixa)})
		case queroSacar > somaNotas(meuCaixa):
			caixaInsuficiente++
			caixaInsufSeguidos++
			extratos = append(extratos, extrato{requisitado: queroSacar, sacado: meuSaque, flag: -1, emCaixa: somaNotas(meuCaixa)})
		default:
			fmt.Println("algum erro")
		}
		if somaNotas(meuCaixa) == 0 || caixaInsufSeguidos == 10 {
			break
		}
	}
	fmt.Println("")
	fmt.Printf("==== FORAM REALIZADOS ====\n")
	fmt.Printf("%v saques\n%v saques rejeitados por falta de notas\n%v solicitações de saque maior que o caixa disponível.\n", saquesOK, saquesParciais, caixaInsuficiente)

	fmt.Println("")
	fmt.Println("==== CAIXA RESTANTE ====")
	contarNotas(meuCaixa)
	fmt.Printf("Total em caixa: $%v\n", somaNotas(meuCaixa))

	fmt.Println("")
	fmt.Println("==== CAIXA DESLIGADO ====")

	fmt.Println("")
	fmt.Println("==== RESUMO DA OPERAÇÃO ====")
	fmt.Println("==== Caixa Inicial ====")
	contarNotas(caixaInicial)
	fmt.Printf("Total inicial em caixa: $%v\n", somaNotas(caixaInicial))
	fmt.Println("==== Operações realizadas ====")
	for i := range extratos {
		fmt.Println(i, extratos[i])
	}
}

type valor int

const nota100 valor = 100
const nota050 valor = 50
const nota020 valor = 20
const nota010 valor = 10
const nota005 valor = 5
const nota002 valor = 2
const nota001 valor = 1

type notas struct {
	face valor
	qtd  int
}

type caixa []notas

func caixaZerado[T caixa | saque]() T {
	var c = T{
		{face: nota100},
		{face: nota050},
		{face: nota020},
		{face: nota010},
		{face: nota005},
		{face: nota002},
		{face: nota001},
	}
	return c
}

func iniciarCaixaAleatorio(nroMaxNotas int) caixa {
	fmt.Println("")
	fmt.Println("==== CAIXA INICIALIZADO ALEATORIAMENTE ====")
	c := caixaZerado[caixa]()
	for i := range c {
		c[i].qtd = int(rand.Intn(nroMaxNotas))
	}
	contarNotas[caixa](c)
	return c
}

func contarNotas[T caixa | saque](c T) {
	if len(c) == 0 {
		fmt.Println("Não há notas a entregar")
	}
	for _, v := range c {
		switch {
		case v.qtd > 1:
			fmt.Printf("%v notas de $%v\n", v.qtd, v.face)
		case v.qtd == 1:
			fmt.Printf("%v nota de $%v\n", v.qtd, v.face)
		default:
			continue
		}
	}
}

func somaNotas[T caixa | saque](c T) valor {
	var soma valor
	for _, v := range c {
		soma += valor(v.qtd) * v.face
	}
	return soma
}

type saque []notas

func requisitarNotas(v valor, c caixa) (saque, caixa) {
	fmt.Printf("valor a sacar: $%v\n", v)
	var s saque
	if v > somaNotas(c) {
		fmt.Println("==== NÃO HÁ DINHEIRO SUFICIENTE ====")
		fmt.Println("===== SOLICITE UM VALOR MENOR! =====")
		return s, c
	}

	var cRestante caixa

	qtd := int(v / c[0].face)
	if c[0].qtd-int(qtd) < 0 {
		qtd = int(c[0].qtd)
	}
	cRestante = append(cRestante, notas{face: c[0].face, qtd: (c[0].qtd - int(qtd))})
	fmt.Printf("entregar %v notas de $%v\n", qtd, c[0].face)
	fmt.Printf("restaram %v notas de $%v no caixa\n", cRestante[0].qtd, cRestante[0].face)

	if len(c) == 1 || v == valor(qtd)*c[0].face {
		s = append(s, notas{face: c[0].face, qtd: int(qtd)})
		return s, c
	}
	s = append(s, notas{face: c[0].face, qtd: int(qtd)})
	sresto, cresto := requisitarNotas(v-valor(qtd)*c[0].face, c[1:])
	s = append(s, sresto...)

	cRestante = append(cRestante, cresto...)

	if somaNotas(s) != v {
		s = caixaZerado[saque]()
		cRestante = c
	}

	return s, cRestante
}

type extrato struct {
	flag        int
	requisitado valor
	emCaixa     valor
	sacado      saque
}

/*
Desenvolva um programa que simule a entrega de notas quando um cliente efetuar um saque em um caixa eletrônico. Os requisitos básicos são os seguintes:

Entregar o menor número de notas;
É possível sacar o valor solicitado com as notas disponíveis;
Saldo do cliente infinito;
Quantidade de notas infinito (pode-se colocar um valor finito de cédulas para aumentar a dificuldade do problema);
Notas disponíveis de R$ 100,00; R$ 50,00; R$ 20,00 e R$ 10,00

https://github.com/golang-cwb/dojos/blob/master/2020-04-23_saque/solution.go

*/
