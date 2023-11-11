package main

import (
	"fmt"
	"time"
)

var publisherID int
var workerID int

func main() {
	c := make(chan string)
	for w := 0; w < 3; w++ {
		go worker(c)
	}
	for p := 0; p < 12; p++ {
		go publisher(c)
	}
	time.Sleep(time.Second)
	fmt.Println("Jornada de trabalho encerrada.")
}

// Publisher is a fan out process
func publisher(out chan string) {
	publisherID++
	thisId := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisId)
		data := fmt.Sprintf("Data from publisher %d. Data: %d", thisId, dataID)
		out <- data
	}
}

// Worker is a fan in process
func worker(in <-chan string) {
	workerID++
	thisId := workerID
	for {
		fmt.Printf("worker %d: waiting for input...\n", thisId)
		input := <-in
		fmt.Printf("worker %d: input is\t%s\n", thisId, input)
	}
}
