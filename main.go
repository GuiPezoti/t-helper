package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
	"github.com/GuiPezoti/t-helper/internal/ui/models"
)

func main() {
	model := models.NewAppModel()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro: %v", err)
        os.Exit(1)
    }
}