package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := initialiseCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		command := cleanInput(text)

		if value, ok := commands[command[0]]; ok {
			value.callback()
		} else {
			fmt.Println("Unknown command")
		}

	}
}
