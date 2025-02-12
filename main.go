package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jacob-cantrell/pokedexcli/internal/pokeapi"
	"github.com/jacob-cantrell/pokedexcli/internal/pokecache"
)

type cli struct {
	commands map[string]cliCommand
}

func (c *cli) getCommands() map[string]cliCommand {
	return c.commands
}

func (c *cli) addCommand(cmd cliCommand) {
	c.commands[cmd.name] = cmd
}

type cliCommand struct {
	name        string
	description string
	callback    func(string, *pokeapi.Config) error
}

func commandExit(s string, con *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (c *cli) commandHelp(s string, con *pokeapi.Config) error {
	// Now you can access c.commands directly
	// Your help logic here
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, v := range c.commands {
		fmt.Printf("\n%s: %s", v.name, v.description)
	}
	fmt.Print("\n")
	return nil
}

func CleanInput(text string) []string {
	s := strings.ToLower(text)
	wordList := strings.Fields(s)
	return wordList
}

var pokeCache *pokecache.Cache

func main() {
	pokeCache = pokecache.NewCache(5 * time.Minute)
	pokedex := make(map[string]pokeapi.Pokemon)

	cliObj := cli{
		commands: make(map[string]cliCommand),
	}
	con := pokeapi.Config{
		Next:      "",
		Previous:  "",
		PokeCache: pokeCache,
		Pokedex:   pokedex,
	}
	conPtr := &con
	cliObj.addCommand(cliCommand{
		name:        "catch",
		description: "Attempts to catch a pokemon!",
		callback:    pokeapi.Catch,
	})
	cliObj.addCommand(cliCommand{
		name:        "explore",
		description: "Lists all pokemon in a given location area",
		callback:    pokeapi.Explore,
	})
	cliObj.addCommand(cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    cliObj.commandHelp,
	})
	cliObj.addCommand(cliCommand{
		name:        "map",
		description: "Shows next 20 locations",
		callback:    pokeapi.Map,
	})
	cliObj.addCommand(cliCommand{
		name:        "mapb",
		description: "Shows previous 20 locations",
		callback:    pokeapi.Mapb,
	})
	cliObj.addCommand(cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	})

	commands := cliObj.getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			commandSplit := CleanInput(text)
			if len(commandSplit) == 0 {
				continue
			}
			if _, ok := commands[commandSplit[0]]; !ok {
				fmt.Println("Unknown command")
			} else {
				if len(commandSplit) > 1 {
					err := commands[commandSplit[0]].callback(commandSplit[1], conPtr)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					err := commands[commandSplit[0]].callback("", conPtr)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}
