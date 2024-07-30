package cmd

import (
	"fmt"
)

func CommandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Descp)
	}
	fmt.Println()
	return nil
}
