package main

import (
	"fmt"
	"os"

	pokeapi "github.com/Mars-weald/Pokedex/internal/pokeAPI"
)

type commandCLI struct {
	name     string
	desc     string
	callback func(c *config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage: \n\n")
	for _, key := range reg {
		fmt.Printf("%s: %s\n", key.name, key.desc)
	}
	return nil
}

func commandMapf(c *config) error {
	areaLocations, err := pokeapi.Pokemap(c.nextURL)
	if err != nil {
		return fmt.Errorf("ERROR getting locations. Check subfile: %w", err)
	}

	c.nextURL = areaLocations.Next
	c.prevURL = areaLocations.Previous

	for _, area := range areaLocations.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.prevURL == nil {
		return fmt.Errorf("You're on the first page\n")
	}

	areaLocations, err := pokeapi.Pokemap(c.prevURL)
	if err != nil {
		return fmt.Errorf("ERROR getting locations. Check subfile: %w", err)
	}

	for _, area := range areaLocations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
