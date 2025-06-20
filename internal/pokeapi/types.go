package pokeapi

type PaginatedLocationAreas struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type LocationArea struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon EncounteredPokemon `json:"pokemon"`
}

type EncounteredPokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	Stat  StatDetail `json:"stat"`
	Value int        `json:"base_stat"`
}

type StatDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonType struct {
	Type PokemonTypeDetail `json:"type"`
}

type PokemonTypeDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
