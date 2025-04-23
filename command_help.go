package main

import "fmt"

func commandHelp(config *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range initialiseCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
