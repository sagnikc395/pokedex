package main

import (
	"fmt"

	"github.com/sagnikc395/pokedex/api"
)

func main() {
	// repl.StartRepl()
	res, err := api.GetPokemonLocationAreas("5")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Name)
	fmt.Println(res.Url)
}
