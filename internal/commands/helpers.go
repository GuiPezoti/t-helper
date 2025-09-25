package commands

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:	"exit",
			description: "Exit t-helper",
			Callback: commandExit,
		},
		"help": {
			name:	"help",
			description: "Show all availble commands",
			Callback: commandHelp,
		},
	}
}