package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	s := strings.ToLower(text)
	wordList := strings.Fields(s)
	return wordList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			text = strings.ToLower(text)
			commandSplit := strings.Fields(text)
			if len(commandSplit) == 0 {
				continue
			}
			fmt.Printf("Your command was: %s\n", commandSplit[0])
		}
	}
}
