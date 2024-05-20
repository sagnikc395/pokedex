package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	var input string
	for {
		fmt.Scanf("%s", &input)

	}
}
