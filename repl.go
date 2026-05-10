package main

import (
	"fmt"
	"strings"
)

func CleanInput(text string) []string {
	var cleanedString []string
	trimString := strings.TrimSpace(text)
	splitString := strings.Split(trimString, "  ") //clean double spacing
	stringed := strings.Join(splitString, " ")     //replaces double spacing with single
	cleaved := strings.Split(stringed, " ")
	for i := range cleaved {
		cleanedString = append(cleanedString, strings.ToLower(cleaved[i]))
	}
	fmt.Println(cleanedString)
	return cleanedString
}
