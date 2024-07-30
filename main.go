package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Command struct {
	Name     string
	Descp    string
	Callback func() error
}

func commandHelp() error {
	fmt.Printf("\nWelcom to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func main() {
	var prompt string = "pokedex > "
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]Command{
		"help": {
			Name:     "help",
			Descp:    "Displays a help message",
			Callback: commandHelp,
		},
		"exit": {
			Name:     "exit",
			Descp:    "Exit the Pokedex",
			Callback: commandExit,
		},
	}

	for {
		fmt.Printf(prompt)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			break
		}

		cmd := scanner.Text()
		k, status := commands[cmd]
		if status == false {
			fmt.Printf("command not found")
		}
		err := k.Callback()
		if err != nil {
			log.Fatal(err)
		}
	}
}
