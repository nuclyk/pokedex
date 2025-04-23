package main

import (
	"fmt"

	"github.com/nuclyk/pokedex/internal/pokeapi"
)

func commandExplore(config *config, arg string) error {
	if arg == "" {
		fmt.Println("Please provide location which you would like to explore.")
		return nil
	}

	fmt.Printf("Exploring %s ...", arg)
	fullURL := baseUrl + arg
	pokemons, err := pokeapi.GetAreaPokemons(fullURL)

	if err != nil {
		fmt.Println("Could not find the area. Please try again.\n")
		return err
	} else {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range pokemons.PokemonEncounters {
			fmt.Printf("- %s\n", pokemon.Pokemon.Name)
		}
	}

	return nil
}
