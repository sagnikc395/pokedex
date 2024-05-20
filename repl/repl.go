package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func execCmdHelp() error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help:  Displays a help message
map:   Displays the names of 20 location areas in the Pokemon world.
mapb:  Displays the names of the 20 previous location areas in the Pokemon world.
exit:  Exit the Pokedex`)
	return nil
}

func execCmdExit() error {
	fmt.Println("Bye!")
	os.Exit(0)
	return nil
}

var cmds = map[string]CliCommand{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    execCmdHelp,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    execCmdExit,
	},
}

func CommandLookup(cmd string) CliCommand {
	switch cmd {
	case "help":
		return CliCommand{
			Name:        "help",
			Description: "Displays a help message",
			Callback:    execCmdHelp,
		}
	case "exit":
		return CliCommand{
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    execCmdExit,
		}
	default:
		return CliCommand{
			Name:        "",
			Description: "",
			Callback:    nil,
		}
	}
}

func StartRepl() {
	fmt.Printf("Pokedex > ")
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := strings.Fields(reader.Text())
		userCmd := input[0]
		if clicmd, ok := cmds[userCmd]; ok {
			err := clicmd.Callback()
			if err != nil {
				fmt.Println("error : ", err)
			}
		} else {
			fmt.Println("Invalid command: use `help` if you're stuck.")
		}
		fmt.Printf("Pokedex > ")
	}

}
