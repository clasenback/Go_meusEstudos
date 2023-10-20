package main

import (
	"fmt"
	"math/rand"
	"time"
)

const JOGADOR1 string = "🏓 ping"
const JOGADOR2 string = "pong 🏓"
const TJOGADA int = 2000
const TMAX time.Duration = 1900 * time.Millisecond

type jogada struct {
	jogada  int
	jogador string
	tempo   time.Duration
}

func main() {
	canal := make(chan jogada)
	comm1 := make(chan bool)
	comm2 := make(chan bool)

	go func() {
		go iniciaPP(comm2)
		go jogadorPP(JOGADOR1, canal, comm2, comm1)
		go jogadorPP(JOGADOR2, canal, comm1, comm2)
	}()

	go assisteJogo(canal)

	var quit string
	fmt.Scanln(&quit)
}

func iniciaPP(bolaEnviada chan<- bool) { bolaEnviada <- true }

func jogadorPP(jogador string, transmite chan<- jogada, bolaRecebida <-chan bool, bolaEnviada chan<- bool) {
	var turno int
	var bola bool
	var evento jogada
	var t time.Duration
	for {
		turno++
		t = time.Duration(rand.Intn(TJOGADA)) * time.Millisecond
		time.Sleep(t)
		bola = <-bolaRecebida
		evento.jogador = jogador
		evento.jogada = turno
		evento.tempo = t
		transmite <- evento
		bolaEnviada <- bola
	}
}

func assisteJogo(canal chan jogada) {
	for v := range canal {
		if v.tempo > TMAX {
			fmt.Printf("%v\t%v\t\tPERDEU na jogada %v: %v maior que %v\n", v.jogada, v.jogador, v.jogada, v.tempo, TMAX)
			fmt.Println()
			fmt.Println("⌨\nDigite ENTER para encerrar o programa...")
			return
		} else {
			fmt.Printf("%v\t%v\t\t%v\n", v.jogada, v.jogador, v.tempo)
		}
	}
}
