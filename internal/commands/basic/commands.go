package basic

import ("os"
		"fmt"
		"flag"
	)

func commandExit(fs *flag.FlagSet, args []string) error {
	force := fs.Bool("force", false, "Force exit without confirmation")
	fs.Parse(args)
	if *force {
		os.Exit(0)
	}
	fmt.Printf("Closing T-helper... See ya!\n")
	os.Exit(0)
	return nil
}

func commandHelp(fs *flag.FlagSet, args []string) error {
	verbose := fs.Bool("verbose", false, "Show detailed information")
	fs.BoolVar(verbose, "v", false, "Show detailed information (shorthand)")

	fs.Parse(args)

	fmt.Println()
	fmt.Println("Welcome to T-helper!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommandsSorted() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
		if *verbose {
			tempFS := flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
			tempFS.SetOutput(os.Stdout)
			fmt.Printf("  Flags: use '%s --help' to see available flags\n", cmd.Name)
		}
	}
	fmt.Println()
	return nil
}