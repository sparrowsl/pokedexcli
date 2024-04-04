package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name provided!")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("Failed to catch %s!\n", pokemonName)
	}

	// Save the caught pokemon
	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught\n", pokemonName)

	return nil
}
