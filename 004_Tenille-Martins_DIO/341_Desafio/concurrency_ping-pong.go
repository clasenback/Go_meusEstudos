package main

import (
	"fmt"
	"math/rand"
	"time"
)

const JOGADOR1 nome = "🏓 ping"
const JOGADOR2 nome = "pong 🏓"
const TJOGADA int = 2000
const TMAX time.Duration = 1900 * time.Millisecond

type atleta int

type nome string

type jogada struct {
	jogada   int
	jogador  nome
	tempo    time.Duration
	contador atleta
}

func main() {
	player1, player2 := criaJogadores()
	partida := make(chan jogada)

	go pingPong(player1, player2, partida)

	go assistePartida(partida)

	var quit string
	fmt.Scanln(&quit)
}

func criaJogadores() (player1, player2 chan atleta) {
	player1 = make(chan atleta)
	player2 = make(chan atleta)
	return
}

func iniciaPP(enviadaPara chan<- atleta) { enviadaPara <- 0 }

func jogadorPP(jogador nome, recebeDe <-chan atleta, enviaPara chan<- atleta, partida chan<- jogada) {
	var turno int
	var bola atleta
	var evento jogada
	var t time.Duration
	for {
		turno++
		t = time.Duration(rand.Intn(TJOGADA)) * time.Millisecond
		time.Sleep(t)
		bola = <-recebeDe
		bola++
		evento.contador = bola
		evento.jogador = jogador
		evento.jogada = turno
		evento.tempo = t
		partida <- evento
		enviaPara <- bola
	}
}

func pingPong(p1, p2 chan atleta, partida chan jogada) {
	go iniciaPP(p2)
	go jogadorPP(JOGADOR1, p2, p1, partida)
	go jogadorPP(JOGADOR2, p1, p2, partida)
}

func assistePartida(partida chan jogada) {
	for evento := range partida {
		if evento.tempo > TMAX {
			fmt.Printf("%v\t%v\t\tPERDEU na jogada %v: %v maior que %v\n", evento.jogada, evento.jogador, evento.jogada, evento.tempo, TMAX)
			fmt.Println("Contador:", evento.contador)
			fmt.Println()
			fmt.Println("⌨\nDigite ENTER para encerrar o programa...")
			return
		} else {
			fmt.Printf("%v\t%v\t\t%v\n", evento.jogada, evento.jogador, evento.tempo)
		}
	}
}
