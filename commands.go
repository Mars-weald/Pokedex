package main

import (
	"fmt"
	"os"
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
	areaLocations := pokemap(c.baseURL)
	return nil
}

func commandMapb(c *config) error {
	return nil
}
