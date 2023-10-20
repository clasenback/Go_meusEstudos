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
	player1, player2 := criaRaquetes()
	partida := make(chan jogada)

	go pingPong(player1, player2, partida)

	go assistePartida(partida)

	var quit string
	fmt.Scanln(&quit)
}

func criaRaquetes() (raquete1, raquete2 chan bool) {
	raquete1 = make(chan bool)
	raquete2 = make(chan bool)
	return
}

func iniciaPP(enviadaPara chan<- bool) { enviadaPara <- true }

func jogadorPP(jogador string, recebeDe <-chan bool, enviaPara chan<- bool, partida chan<- jogada) {
	var turno int
	var bola bool
	var evento jogada
	var t time.Duration
	for {
		turno++
		t = time.Duration(rand.Intn(TJOGADA)) * time.Millisecond
		time.Sleep(t)
		bola = <-recebeDe
		evento.jogador = jogador
		evento.jogada = turno
		evento.tempo = t
		partida <- evento
		enviaPara <- bola
	}
}

func pingPong(p1, p2 chan bool, partida chan jogada) {
	go iniciaPP(p2)
	go jogadorPP(JOGADOR1, p2, p1, partida)
	go jogadorPP(JOGADOR2, p1, p2, partida)
}

func assistePartida(partida chan jogada) {
	for evento := range partida {
		if evento.tempo > TMAX {
			fmt.Printf("%v\t%v\t\tPERDEU na jogada %v: %v maior que %v\n", evento.jogada, evento.jogador, evento.jogada, evento.tempo, TMAX)
			fmt.Println()
			fmt.Println("⌨\nDigite ENTER para encerrar o programa...")
			return
		} else {
			fmt.Printf("%v\t%v\t\t%v\n", evento.jogada, evento.jogador, evento.tempo)
		}
	}
}
