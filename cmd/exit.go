package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var exitCmd = &cobra.Command{
	Use:     "exit",
	Aliases: []string{"quit"},
	Short:   "Exit t-helper",
	Long:    `Exit the interactive mode.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("__EXIT__")
	},
}

func init() {
	rootCmd.AddCommand(exitCmd)
}