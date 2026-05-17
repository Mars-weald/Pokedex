package main

import (
	"bufio"
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

		"mapf": {
			name:     "mapf",
			desc:     "Displays the next page of location areas",
			callback: commandMapf,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Displays the previous page of location areas",
			callback: commandMapb,
		},
	}
}

type config struct {
	baseURL string
	nextURL string
	prevURL string
}

func startoop(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; ; i++ {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := strings.ToLower(input)
		words := strings.Split(cleanedInput, " ")
		command := words[0]

		_, exists := reg[command]
		if exists {
			err := reg[command].callback(conf)
			if err != nil {
				fmt.Printf("%v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
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
