package pokeapi

import "fmt"

func commandPokedex(config *Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range config.Pokemon {
		fmt.Printf(" - %s\n", p.Name)
	}

	return nil
}
