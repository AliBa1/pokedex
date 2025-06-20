package pokeapi

import "fmt"

func commandHelp(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, c := range GetCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	// fmt.Println("help: Displays a help message")
	// fmt.Println("exit: Exit the Pokedex")
	return nil
}
