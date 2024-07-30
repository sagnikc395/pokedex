package cmd

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:     "help",
			Descp:    "Displays a help message",
			Callback: CommandHelp,
		},
		"exit": {
			Name:     "exit",
			Descp:    "Exit the Pokedex",
			Callback: CommandExit,
		},
		"map": {
			Name:     "map",
			Descp:    "list locations",
			Callback: CommandMakeAPIRequest,
		},
	}
}
