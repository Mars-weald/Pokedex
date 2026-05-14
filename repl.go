package main

import (
	"fmt"
	"os"
	"strings"
)

var reg = make(map[string]commandCLI)

func init() {
	reg = map[string]commandCLI{
		"exit": {
			name:     "exit",
			desc:     "Exits the Pokedex",
			callback: commandExit,
		},

		"help": {
			name:     "help",
			desc:     "Displays help message",
			callback: commandHelp,
		},
	}
}

func CleanInput(text string) []string {
	var cleanedString []string
	trimString := strings.TrimSpace(text)
	splitString := strings.Split(trimString, "  ") //clean double spacing
	stringed := strings.Join(splitString, " ")     //replaces double spacing with single
	cleaved := strings.Split(stringed, " ")
	for i := range cleaved {
		cleanedString = append(cleanedString, strings.ToLower(cleaved[i]))
	}
	return cleanedString
}

type commandCLI struct {
	name     string
	desc     string
	callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {

	fmt.Print("Welcome to the Pokedex!\nUsage: \n\n")
	for _, key := range reg {
		fmt.Printf("%s: %s\n", key.name, key.desc)
	}
	return nil
}
