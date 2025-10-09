package cmd

import (
	"fmt"
	"github.com/GuiPezoti/t-helper/internal/app"
	"github.com/charmbracelet/lipgloss" // <-- Adicione esta importação
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show available commands",
	Long:  `Display all available commands and their descriptions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.TitleStyle.Render("Available commands:"))
		fmt.Println() // Adiciona uma linha em branco para espaçamento

		commands := []struct {
			Name string
			Desc string
		}{
			{"clear", "Clear the screen"},
			{"exit", "Exit t-helper"},
			{"help", "Show available commands"},
		}

		// Calcula o tamanho máximo do nome do comando para alinhar a coluna
		maxLen := 0
		for _, c := range commands {
			if len(c.Name) > maxLen {
				maxLen = len(c.Name)
			}
		}

		// Itera e usa lipgloss para criar as linhas formatadas
		for _, c := range commands {
			// Cria uma "célula" para o nome do comando com largura fixa
			// Isso garante que todas as descrições comecem na mesma posição
			nameCell := lipgloss.NewStyle().
				Width(maxLen). // Define a largura baseada no comando mais longo
				Render(app.CommandStyle.Render(c.Name))

			// Junta o nome do comando, um separador, e a descrição horizontalmente
			line := lipgloss.JoinHorizontal(lipgloss.Left,
				nameCell,
				"  - ",
				app.DescriptionStyle.Render(c.Desc),
			)

			fmt.Println(line)
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}