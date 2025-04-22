package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (locationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var areas locationAreas
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}

	return areas, nil
}
