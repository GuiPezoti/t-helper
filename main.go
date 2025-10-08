// main.go
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github/GuiPezoti/t-helper/internal/app"
)

func main() {
	// Cria o programa Bubbletea
	p := tea.NewProgram(
		app.NewModel(),           // Estado inicial
		tea.WithAltScreen(),      // Usa tela alternativa (limpa ao sair)
		tea.WithMouseCellMotion(), // Habilita mouse
	)

	// Executa o programa
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(1)
	}
}
