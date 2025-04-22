package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, command := range initialiseCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
