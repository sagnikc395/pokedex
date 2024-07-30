package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sagnikc395/pokedex/cmd"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func Repl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		cmd, exists := cmd.GetCommands()[cmdName]
		if exists {
			err := cmd.Callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}
