package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Next     string
	Previous string
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Map(con *Config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	var url string
	if con.Previous == "" && con.Next == "" { // Original URL
		url = baseURL
	} else if con.Next == "" { // Last page
		fmt.Println("You are on the last page!")
		return nil
	} else { // Use con.Next for next page
		url = con.Next
	}

	// Get Response and check if valid
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	// Read body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Unmarshal
	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	// Print locations
	for _, result := range locations.Results {
		fmt.Printf("%s\n", result.Name)
	}

	// Update Next and Previous
	con.Next = locations.Next
	if locations.Previous == nil {
		con.Previous = ""
	} else {
		con.Previous = *locations.Previous
	}

	return nil
}

func Mapb(con *Config) error {
	var url string
	if con.Previous == "" && con.Next == "" { // First Page
		fmt.Println("you're on the first page")
		return nil
	} else { // All pages outside of First page has a previous page
		url = con.Previous
	}

	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	// Get Response and check if valid
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	// Read body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Unmarshal
	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return err
	}

	// Print locations
	for _, result := range locations.Results {
		fmt.Printf("%s\n", result.Name)
	}

	// Update Next and Previous
	con.Next = locations.Next
	if locations.Previous == nil {
		con.Previous = ""
	} else {
		con.Previous = *locations.Previous
	}

	return nil
}
