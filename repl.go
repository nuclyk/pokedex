package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

type config struct {
	Next     string
	Previous string
}

func startRepl() {
	config := config{
		Next:     baseUrl,
		Previous: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	commands := initialiseCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		if cmd, ok := commands[words[0]]; ok {
			err := cmd.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func initialiseCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show area locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous area locations",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}
