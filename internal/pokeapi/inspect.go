package pokeapi

import "fmt"

func commandInspect(config *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("specify the name of the pokemon you want to inspect")
	}
	selectedPokemonName := args[0]

	pokemon, ok := config.Pokemon[selectedPokemonName]
	if !ok {
		return fmt.Errorf("you have not caught %s", pokemon.Name)
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", s.Stat.Name, s.Value)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("	-%s\n", t.Type.Name)
	}

	return nil
}
