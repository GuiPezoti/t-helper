// internal/app/commands.go
package app

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
)

// GetCommands retorna todos os comandos disponíveis
func GetCommands() []Command {
	return []Command{
		{
			Name:        "help",
			Description: "Show available commands",
			Execute:     executeHelp,
		},
		{
			Name:        "exit",
			Description: "Exit t-helper",
			Execute:     executeExit,
		},
		{
			Name:        "clear",
			Description: "Clear the screen",
			Execute:     executeClear,
		},
	}
}

// executeHelp mostra todos os comandos disponíveis
func executeHelp(m *Model, args []string) tea.Cmd {
	m.output = append(m.output, "")
	m.output = append(m.output, "Available commands:")
	m.output = append(m.output, "")
	
	// Ordena comandos alfabeticamente
	commands := make([]Command, len(m.commands))
	copy(commands, m.commands)
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})
	
	// Lista comandos
	for _, cmd := range commands {
		line := fmt.Sprintf("  %-10s - %s", cmd.Name, cmd.Description)
		m.output = append(m.output, line)
	}
	
	m.output = append(m.output, "")
	
	return nil
}

// executeExit sai da aplicação
func executeExit(m *Model, args []string) tea.Cmd {
	m.output = append(m.output, "Goodbye! 👋")
	
	// Retorna comando para sair
	return tea.Quit
}

// executeClear limpa o histórico
func executeClear(m *Model, args []string) tea.Cmd {
	m.output = []string{}
	return nil
}
