// internal/app/model.go
package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model é o estado da aplicação
// Pense nele como "o que a aplicação sabe agora"
type Model struct {
	input       string   // O que o usuário está digitando
	commands    []Command // Comandos disponíveis
	output      []string  // Histórico de saídas
	cursorPos   int       // Posição do cursor no input
	ready       bool      // App está pronto?
}

// Command representa um comando disponível
type Command struct {
	Name        string
	Description string
	Execute     func(*Model, []string) tea.Cmd
}

// NewModel cria o estado inicial da aplicação
func NewModel() Model {
	return Model{
		input:    "",
		commands: GetCommands(), // Carrega comandos disponíveis
		output:   []string{"Welcome to T-helper! Type 'help' for commands."},
		ready:    true,
	}
}

// Init é chamado uma vez no início
// Retorna um comando inicial (ou nil)
func (m Model) Init() tea.Cmd {
	return nil
}
