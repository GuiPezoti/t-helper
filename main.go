// main.go
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/GuiPezoti/t-helper/cmd"
	"github.com/GuiPezoti/t-helper/internal/app"
)

func main() {
	// Sempre inicia no modo interativo com BubbleTea
	// mas passa os comandos Cobra para execução
	p := tea.NewProgram(
		app.NewModel(cmd.GetRootCmd()),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(1)
	}
}