package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchData(url string, config *Config) (data []byte, error error) {
	data, inCache := config.cache.Get(url)
	if inCache {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error during location area GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading location area body: %v", err)
	}

	config.cache.Add(url, body)

	return body, nil
}

func fetchPokemon(url string, config *Config) (Pokemon, error) {
	var pokemon Pokemon

	data, err := fetchData(url, config)
	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("error unmarshaling pokemon: %v", err)
	}

	return pokemon, nil
}

func fetchLocationArea(url string, config *Config) (LocationArea, error) {
	var locationArea LocationArea

	data, err := fetchData(url, config)
	if err != nil {
		return locationArea, err
	}

	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return locationArea, fmt.Errorf("error unmarshaling location area: %v", err)
	}

	return locationArea, nil
}

func fetchPaginatedLocationAreas(url string, config *Config) (PaginatedLocationAreas, error) {
	var paginatedLocationAreas PaginatedLocationAreas

	data, err := fetchData(url, config)
	if err != nil {
		return paginatedLocationAreas, err
	}

	err = json.Unmarshal(data, &paginatedLocationAreas)
	if err != nil {
		return paginatedLocationAreas, fmt.Errorf("error unmarshaling paginated response for location area: %v", err)
	}

	return paginatedLocationAreas, nil
}
