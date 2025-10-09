package app

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Estilos usando Lipgloss
var (
	promptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)
	
	inputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86"))
	
	outputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
	
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("170")).
		Bold(true).
		Padding(0, 1)
)

// View renderiza o estado atual na tela
func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}
	
	var b strings.Builder
	
	// Título
	b.WriteString(titleStyle.Render("T-HELPER"))
	b.WriteString("\n\n")
	
	// Output (últimas 10 linhas)
	startIdx := 0
	if len(m.output) > 10 {
		startIdx = len(m.output) - 10
	}
	
	for i := startIdx; i < len(m.output); i++ {
		b.WriteString(outputStyle.Render(m.output[i]))
		b.WriteString("\n")
	}
	
	b.WriteString("\n")
	
	// Prompt + Input
	b.WriteString(promptStyle.Render("T-helper > "))
	b.WriteString(inputStyle.Render(m.input))
	
	// Cursor
	if len(m.input) == m.cursorPos {
		b.WriteString(inputStyle.Render("_"))
	}
	
	b.WriteString("\n\n")
	
	// Ajuda no rodapé
	b.WriteString(outputStyle.Render("Press Ctrl+C to quit | Type 'help' for commands"))
	
	return b.String()
}