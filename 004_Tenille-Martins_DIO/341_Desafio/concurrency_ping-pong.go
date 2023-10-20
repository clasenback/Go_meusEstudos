package main

import (
	"fmt"
	"math/rand"
	"time"
)

type pacote struct {
	pp    string
	n     int
	tempo time.Duration
}

func main() {
	canal := make(chan pacote)
	comm1 := make(chan bool)
	comm2 := make(chan bool)

	go func() {
		go iniciaPP(comm1)
		go pingpong("ping", canal, comm1, comm2)
		go pingpong("pong", canal, comm2, comm1)
	}()

	go func() {
		for v := range canal {
			fmt.Println(v.n, v.pp, v.tempo)
		}
	}()

	var quit string
	fmt.Scanln(&quit)

}

func pingpong(s string, c chan<- pacote, p1 <-chan bool, p2 chan<- bool) {
	var i int
	var bastao bool
	var pct pacote
	var t time.Duration
	for {
		t = time.Duration(rand.Intn(2000) * int(time.Millisecond))
		time.Sleep(t)
		i++
		bastao = <-p1
		pct.pp = s
		pct.n = i
		pct.tempo = t
		c <- pct
		p2 <- bastao
	}
}

func iniciaPP(p chan<- bool) { p <- true }
