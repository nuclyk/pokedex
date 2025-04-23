package main

import (
	"fmt"
	"log"

	"github.com/nuclyk/pokedex/internal/pokeapi"
)

func commandMapf(config *config) error {

	areas, err := pokeapi.GetLocationAreas(config.Next)
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
