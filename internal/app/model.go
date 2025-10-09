package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// Model é o estado da aplicação (apenas UI)
type Model struct {
	input     string         // O que o usuário está digitando
	rootCmd   *cobra.Command // Comandos Cobra
	output    []string       // Histórico de saídas
	cursorPos int            // Posição do cursor no input
	ready     bool           // App está pronto?
}

// NewModel cria o estado inicial da aplicação
func NewModel(rootCmd *cobra.Command) Model {
	return Model{
		input:   "",
		rootCmd: rootCmd,
		output:  []string{rootCmd.Long}, // Usa a mensagem do Cobra
		ready:   true,
	}
}

// Init é chamado uma vez no início
func (m Model) Init() tea.Cmd {
	return nil
}