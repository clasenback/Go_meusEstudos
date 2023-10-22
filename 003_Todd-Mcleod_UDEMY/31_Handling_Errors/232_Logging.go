package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("alguma merda aconteceu! fmt.Println")
		log.Println("alguma merda aconteceu! log.Println")
		log.Panicln("Alguma merda aconteceu! log.Panicln")
		log.Fatalln("Alguma merda aconteceu! log.Fatalln")
	}
}

func foo() {
	log.Println("Alguma merda aconteceu! 💩")
}
