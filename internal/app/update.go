// internal/app/update.go
package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Update é chamado sempre que algo acontece (tecla pressionada, etc)
// Recebe uma mensagem (msg) e retorna o novo estado (Model) e um comando (Cmd)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	// Mensagens de tecla pressionada
	case tea.KeyMsg:
		switch msg.Type {
		
		// Ctrl+C ou Ctrl+D = sair
		case tea.KeyCtrlC, tea.KeyCtrlD:
			return m, tea.Quit
		
		// Enter = executar comando
		case tea.KeyEnter:
			return m.handleCommand()
		
		// Backspace = apagar caractere
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		
		// Seta esquerda
		case tea.KeyLeft:
			if m.cursorPos > 0 {
				m.cursorPos--
			}
		
		// Seta direita
		case tea.KeyRight:
			if m.cursorPos < len(m.input) {
				m.cursorPos++
			}
		
		// Qualquer outra tecla = adicionar ao input
		default:
			// Ignora teclas especiais
			if msg.Type == tea.KeyRunes {
				m.input += string(msg.Runes)
				m.cursorPos = len(m.input)
			}
		}
	
	// Redimensionamento da janela
	case tea.WindowSizeMsg:
		// Pode usar msg.Width e msg.Height se precisar
		return m, nil
	}
	
	return m, nil
}

// handleCommand processa o comando digitado
func (m Model) handleCommand() (tea.Model, tea.Cmd) {
	// Limpa e separa o input em palavras
	input := strings.TrimSpace(m.input)
	if input == "" {
		return m, nil
	}
	
	words := strings.Fields(input)
	cmdName := words[0]
	args := []string{}
	if len(words) > 1 {
		args = words[1:]
	}
	
	// Adiciona comando ao output
	m.output = append(m.output, "$ "+input)
	
	// Procura o comando
	var foundCmd *Command
	for i := range m.commands {
		if m.commands[i].Name == cmdName {
			foundCmd = &m.commands[i]
			break
		}
	}
	
	// Executa ou mostra erro
	var cmd tea.Cmd
	if foundCmd != nil && foundCmd.Execute != nil {
		cmd = foundCmd.Execute(&m, args)
	} else {
		m.output = append(m.output, "Unknown command: "+cmdName)
		m.output = append(m.output, "Type 'help' to see available commands")
	}
	
	// Limpa o input
	m.input = ""
	m.cursorPos = 0
	
	return m, cmd
}