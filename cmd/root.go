package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "t-helper",
	Short: "T-helper - Your terminal helper",
	Long: `Welcome to T-helper! Making the development life a little easier`,
}

// GetRootCmd retorna o comando raiz para ser usado pelo BubbleTea
func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	// Desabilita comandos padrão do Cobra
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}