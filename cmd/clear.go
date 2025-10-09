package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the screen",
	Long:  `Clear the terminal screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("__CLEAR_SCREEN__")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}