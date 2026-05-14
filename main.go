package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
			err := reg[command].callback()
			if err != nil {
				fmt.Printf("%v", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
