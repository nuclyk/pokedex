package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas() {
	fullUrl := baseUrl + "location-area/"

	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err = io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var areas locationAreas
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(locationAreas)

}
