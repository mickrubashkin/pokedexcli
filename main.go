package main

import (
	"time"

	"github.com/mickrubashkin/pokedexcli/internal/pokeapi"
)

func main() {
	pokeCLient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeCLient,
	}
	startRepl(cfg)
}
