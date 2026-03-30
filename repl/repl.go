package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCMD map[string]cliCommand = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    cmdExit,
	},
	"help": {
		name:        "help",
		description: "Help Function of the Pokedex",
		callback:    cmdHelp,
	},
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		cmd, ok := supportedCMD[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}
}
func CleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
