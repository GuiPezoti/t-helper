package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
		Use:   "t-helper",
		Short: "Making the development life a little easier",
		Long:  `T-helper is a CLI tool to help automate development tasks.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to T-helper!")
		},
	}

// Execute executa o comando raiz
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Desabilita comandos padrão do Cobra
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}