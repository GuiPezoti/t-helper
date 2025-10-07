package commands

import "flag"

type BaseCommand struct {
	Name        string
	Description string
	Subcommands []string
	Callback 	func(*flag.FlagSet, []string) error
}