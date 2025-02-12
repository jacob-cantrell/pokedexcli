package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (c *cli) commandHelp() error {
	// Now you can access c.commands directly
	// Your help logic here
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for _, v := range c.commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func CleanInput(text string) []string {
	s := strings.ToLower(text)
	wordList := strings.Fields(s)
	return wordList
}

func main() {
	cliObj := cli{
		commands: make(map[string]cliCommand),
	}
	cliObj.addCommand(cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    cliObj.commandHelp,
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
				err := commands[commandSplit[0]].callback()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
