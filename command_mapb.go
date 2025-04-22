package main

import (
	"fmt"
	"log"

	"github.com/nuclyk/pokedex/internal/pokeapi"
)

func commandMapb(config *config) error {

	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, err := pokeapi.GetLocationAreas(config.Previous)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	config.Next = areas.Next
	config.Previous = areas.Previous

	return nil
}
