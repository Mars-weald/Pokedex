package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const bassURL = "https://pokeapi.co/api/v2/"

type locationList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Pokemap(pageURL *string) (locationList, error) {
	url := bassURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	listOLocations := locationList{}

	res, err := http.Get(url)
	if err != nil {
		return listOLocations, fmt.Errorf("ERROR getting: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return listOLocations, fmt.Errorf("ERROR reading: %w", err)
	}

	err = json.Unmarshal(body, &listOLocations)
	if err != nil {
		return listOLocations, fmt.Errorf("ERROR unmarshalling: %w", err)
	}
	return listOLocations, nil
}
