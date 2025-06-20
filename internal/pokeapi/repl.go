package pokeapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AliBa1/pokedex/internal/pokecache"
)

func Run(cache *pokecache.Cache) {
	config := Config{
		cache:   cache,
		Pokemon: make(map[string]Pokemon),
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Println(cleanedInput)
		command, ok := GetCommands()[cleanedInput[0]]
		args := cleanedInput[1:]
		if ok {
			err := command.callback(&config, args)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	// word := ""
	// cleaned := make([]string, 0)
	// for i, letter := range text {
	// 	if letter != ' ' {
	// 		word += strings.ToLower(string(letter))
	// 	}
	//
	// 	if (letter == ' ' && len(word) > 0) || i+1 == len(text) {
	// 		cleaned = append(cleaned, word)
	// 		word = ""
	// 	}
	// }
	//
	// return cleaned

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
