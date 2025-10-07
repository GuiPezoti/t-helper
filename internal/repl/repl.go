package repl

import (
    "bufio"
	"os"
	"fmt"
	"flag"
	"github/GuiPezoti/t-helper/internal/commands/basic"
)

func StartRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("T-helper > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue	
		}
		command, ok := basic.GetCommands()[words[0]]
		if ok {
			fs := flag.NewFlagSet(command.Name, flag.ContinueOnError)
			fs.SetOutput(os.Stdout)
			args := []string{}
			if len(words) > 1 {
				args = words[1:]
			}
			err := command.Callback(fs, args)
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