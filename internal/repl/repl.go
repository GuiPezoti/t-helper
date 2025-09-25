package repl

import (
    "bufio"
	"strings"
	"os"
	"fmt"
	"github/GuiPezoti/t-helper/internal/commands"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func StartRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("T-helper > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue	
		}
		command, ok := commands.GetCommands()[words[0]]
		if ok {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}