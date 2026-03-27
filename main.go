package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		text = strings.ToLower(text)
		words := strings.Fields(text)
		fmt.Println("Your command was:", words[0])
	}
}
