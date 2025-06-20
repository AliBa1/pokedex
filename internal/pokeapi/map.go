package pokeapi

import (
	"fmt"
)

func printLocations(locations []Location) {
	for _, l := range locations {
		fmt.Println(l.Name)
	}
}

func commandMap(config *Config, args []string) error {
	var url string
	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	} else {
		url = config.Next
	}

	paginatedLocationAreas, err := fetchPaginatedLocationAreas(url, config)
	if err != nil {
		return err
	}

	config.Next = paginatedLocationAreas.Next
	config.Previous = paginatedLocationAreas.Previous
	printLocations(paginatedLocationAreas.Results)

	return nil
}

func commandMapBack(config *Config, args []string) error {
	if config.Previous == "" {
		return fmt.Errorf("cannot map back any further")
	}

	url := config.Previous
	paginatedLocationAreas, err := fetchPaginatedLocationAreas(url, config)
	if err != nil {
		return err
	}

	config.Next = paginatedLocationAreas.Next
	config.Previous = paginatedLocationAreas.Previous
	printLocations(paginatedLocationAreas.Results)

	return nil
}
