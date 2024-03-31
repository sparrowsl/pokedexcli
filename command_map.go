package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.prevLocationAreaURL = resp.Previous
	cfg.nextLocationAreaURL = resp.Next

	return nil
}

func callbackMapb(cfg *config) error {
	// display error if no api can't get previous page
	if cfg.prevLocationAreaURL == nil {
		return errors.New("You're on the first page!!")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.prevLocationAreaURL = resp.Previous
	cfg.nextLocationAreaURL = resp.Next

	return nil
}
