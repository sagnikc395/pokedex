package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BASE_API = "https://pokeapi.co/api/v2"

type LocRespObj struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokemonLocationAreas(id string) (LocRespObj, error) {
	client := &http.Client{}

	url := fmt.Sprintf("%v/location/%v", BASE_API, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error making requests", err)
		return LocRespObj{}, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return LocRespObj{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: API returned status code", resp.StatusCode)
		return LocRespObj{}, err
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return LocRespObj{}, err
	}

	var data LocRespObj
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocRespObj{}, err
	}

	return data, nil

}
