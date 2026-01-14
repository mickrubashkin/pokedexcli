package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(c *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}

func commandMapForward(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Should provide location-area name")
	}

	areaName := args[0]
	fmt.Printf("Exploring %v...\n", areaName)
	locationAreaDetail, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")

	for _, pokemonEncounter := range locationAreaDetail.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Should provide pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("pokemon %s, base experience: %d\n", pokemon.Name, pokemon.BaseExperience)

	result := rand.Intn(pokemon.BaseExperience)
	if result > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Should provide pokemon name")
	}

	name := args[0]
	pokemon, caught := cfg.pokedex[name]
	if caught {
		stats := map[string]int{}
		for _, stat := range pokemon.Stats {
			switch stat.Stat.Name {
			case "hp":
				stats["hp"] = stat.BaseStat
			case "attack":
				stats["attack"] = stat.BaseStat
			case "defense":
				stats["defense"] = stat.BaseStat
			case "special-attack":
				stats["special-attack"] = stat.BaseStat
			case "special-defense":
				stats["special-defense"] = stat.BaseStat
			case "speed":
				stats["speed"] = stat.BaseStat
			}
		}

		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		// fmt.Printf("  -hp: %v\n", stats["hp"])
		// fmt.Printf("  -attack: %v\n", stats["attack"])
		// fmt.Printf("  -defense: %v\n", stats["defense"])
		// fmt.Printf("  -special-attack: %v\n", stats["special-attack"])
		// fmt.Printf("  -special-defense: %v\n", stats["special-defense"])
		// fmt.Printf("  -speed: %v\n", stats["speed"])
		fmt.Println("Types:")

		for _, typeInfo := range pokemon.Types {
			fmt.Printf("  - %v\n", typeInfo.Type.Name)
		}
		return nil
	}

	fmt.Println("you have not caught that pokemon")
	return nil
}
