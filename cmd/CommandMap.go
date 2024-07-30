package cmd

import (
	"io"
	"log"
	"net/http"
)

const BASE_URL = "https://pokeapi.co/api/v2/location/"

func CommandMakeAPIRequest() error {
	resp, err := http.Get(BASE_URL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	sb := string(body)

	locationLists := make([]string, 20)

	return nil
}
