package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Pokedex struct {
	Descriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"descriptions"`
	ID           int    `json:"id"`
	IsMainSeries bool   `json:"is_main_series"`
	Name         string `json:"name"`
	Names        []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEntries []struct {
		EntryNumber    int `json:"entry_number"`
		PokemonSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_entries"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	VersionGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_groups"`
}

func main() {

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/") //mapeia
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Println(response)

	responseData, err := io.ReadAll(response.Body) //dados --> bytes
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))

	var pokedex Pokedex
	json.Unmarshal(responseData, &pokedex)

	//fmt.Println(responseObject.name)
	//fmt.Println((responseObject.Pokemon))

	for n, pokemon := range pokedex.PokemonEntries {
		fmt.Println(n, pokemon.EntryNumber, pokemon.PokemonSpecies.Name)
		fmt.Println("\t", pokemon.PokemonSpecies.URL)
	}
	fmt.Println()

	for n, desc := range pokedex.Descriptions {
		fmt.Println(n, desc.Description)
		fmt.Println("\t", desc.Language.Name, "\t", desc.Language.URL)
	}
	fmt.Println()

	for n, name := range pokedex.Names {
		fmt.Println(n, name.Name)
		fmt.Println("\t", name.Language.Name, "\t", name.Language.URL)
	}
	fmt.Println()

	fmt.Println(pokedex.Region.Name, pokedex.Region.URL)
	fmt.Println()

	for n, v := range pokedex.VersionGroups {
		fmt.Println(n, v.Name, v.URL)
	}
}

/*
type pokedex criado com o auxílio da ferramenta disponível em
https://mholt.github.io/json-to-go/
*/
