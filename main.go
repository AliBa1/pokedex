package main

import (
	"time"

	"github.com/AliBa1/pokedex/internal/pokeapi"
	"github.com/AliBa1/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)

	pokeapi.Run(cache)
}
