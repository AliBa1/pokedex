package pokeapi

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args []string) error {
	if len(args) < 1 {
		// fmt.Println("Include a pokemon when using the catch command")
		// fmt.Println("For example: pokemon pikachu")
		// return nil

		return fmt.Errorf("must include the name of the pokemon your trying to catch")
	}
	selectedPokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", selectedPokemonName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", selectedPokemonName)
	pokemon, err := fetchPokemon(url, config)
	if err != nil {
		return err
	}

	err = catchPokemon(pokemon, config)
	if err != nil {
		return err
	}

	return nil
}

func catchPokemon(pokemon Pokemon, config *Config) error {
	minBaseExperience, maxBaseExperience := 20, 400
	catchProbability := 1.0 - ((float64(pokemon.BaseExperience) - float64(minBaseExperience)) / (float64(maxBaseExperience) - float64(minBaseExperience)))
	randomProbability := rand.Float64()

	caught := randomProbability < catchProbability
	if caught {
		config.Pokemon[pokemon.Name] = pokemon
		fmt.Println(pokemon.Name, "was caught!")
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}
