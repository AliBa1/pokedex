package pokeapi

import (
	"fmt"
)

func printLocationAreaPokemon(locationArea LocationArea) {
	for _, e := range locationArea.Encounters {
		fmt.Println(e.Pokemon.Name)
	}
}

func commandExplore(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("must include the location you want to explore")
	}

	location := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)
	locationArea, err := fetchLocationArea(url, config)
	if err != nil {
		return err
	}

	printLocationAreaPokemon(locationArea)

	return nil
}
