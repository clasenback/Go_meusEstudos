package main

import (
	"fmt"
	"sort"
)

type agente struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

/*
sort the []user by
● age
● last
Also sort each []string “Sayings” for each user
● print everything in a way that is pleasant
*/

type byAge []agente

func (u byAge) Len() int           { return len(u) }
func (u byAge) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u byAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type byLast []agente

func (u byLast) Len() int           { return len(u) }
func (u byLast) Less(i, j int) bool { return u[i].Last < u[j].Last }
func (u byLast) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type bySaying []string

func (u bySaying) Len() int           { return len(u) }
func (u bySaying) Less(i, j int) bool { return u[i] < u[j] }
func (u bySaying) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

func main() {
	u1 := agente{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := agente{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := agente{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	agentes := []agente{u1, u2, u3}

	fmt.Println(agentes)

	// your code goes here
	fmt.Println("===== ===== ===== ===== ===== =====")
	fmt.Println("===== === SORT POR IDADE ==== =====")
	fmt.Println("===== ===== ===== ===== ===== =====")
	sort.Sort(byAge(agentes))
	printaBonito(agentes)
	fmt.Println("===== ===== ===== ===== ===== =====")
	fmt.Println("===== ==== SORT POR NOME ==== =====")
	fmt.Println("===== ===== ===== ===== ===== =====")
	sort.Sort(byLast(agentes))
	printaBonito(agentes)
	fmt.Println("===== ===== ===== ===== ===== =====")
	fmt.Println("===== ==== SORT POR FALA ==== =====")
	fmt.Println("===== ===== ===== ===== ===== =====")
	for a := range agentes {
		sort.Sort(bySaying(agentes[a].Sayings))
	}
	printaBonito(agentes)
}

func printaBonito(agentes []agente) {
	for i := range agentes {
		fmt.Println(agentes[i].First, agentes[i].Last, agentes[i].Age)
		for j := range agentes[i].Sayings {
			fmt.Println("\t", agentes[i].Sayings[j])
		}
	}
}
