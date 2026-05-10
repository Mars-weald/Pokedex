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
		fmt.Printf("Your command was: %s \n", words[0])
	}
}
