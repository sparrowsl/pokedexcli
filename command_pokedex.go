package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, _ ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		return errors.New("No pokemon has been caught")
	}

	fmt.Println("Pokemon in pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
