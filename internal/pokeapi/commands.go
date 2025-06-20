package pokeapi

import (
	"github.com/AliBa1/pokedex/internal/pokecache"
)

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location area names in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch Pokemon to add them to your collection",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "See details about a Pokemon that you've caought",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all Pokemon that you've caught",
			callback:    commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
}

type Config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
	Pokemon  map[string]Pokemon
}
