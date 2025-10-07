package basic

import (
	"sort"
	"github/GuiPezoti/t-helper/internal/commands"
	)
func GetCommands() map[string]commands.BaseCommand {
	return map[string]commands.BaseCommand {
		"exit": {
				Name:	"exit",
				Description: "Exit t-helper",
				Subcommands: nil,
				Callback: commandExit,
			},
		"help": {
				Name:	"help",
				Description: "Show all availble commands",
				Subcommands: nil,
				Callback: commandHelp,
			},
		// "config": {
		// 		Name:	"config",
		// 		Description: "T-helper configuration",
		// 		Subcommands: []string{"get", "set", "list", "delete"},
		// 		Callback: commandHelp,
		// 	},
		"git": {
				Name:	"git",
				Description: "Git commands for faster interaction and development",
				Subcommands: []string{"history"},
				Callback: nil,
			},
		}
}

func GetCommandsSorted() []commands.BaseCommand {
    cmdMap := GetCommands()
    
    names := make([]string, 0, len(cmdMap))
    for name := range cmdMap {
        names = append(names, name)
    }
    sort.Strings(names)
    
    result := make([]commands.BaseCommand, 0, len(cmdMap))
    for _, name := range names {
        result = append(result, cmdMap[name])
    }
    
    return result
}