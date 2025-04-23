package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/nuclyk/pokedex/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2/"
const interval = time.Second * 5

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreas, error) {

	cache := pokecache.NewCache(interval)
	entry, ok := cache.Get(url)
	var areas LocationAreas

	if ok {
		err := json.Unmarshal(entry, &areas)
		if err != nil {
			log.Fatal(err)
		}
	} else {
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

		err = json.Unmarshal(body, &areas)
		if err != nil {
			log.Fatal(err)
		}
	}

	return areas, nil
}
