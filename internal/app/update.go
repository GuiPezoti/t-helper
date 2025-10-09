package app

import (
	"bytes"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Update é chamado sempre que algo acontece (apenas lógica de UI)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
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
			if msg.Type == tea.KeyRunes {
				m.input += string(msg.Runes)
				m.cursorPos = len(m.input)
			}
		}
	
	case tea.WindowSizeMsg:
		return m, nil
	}
	
	return m, nil
}

// handleCommand processa o comando usando Cobra
func (m Model) handleCommand() (tea.Model, tea.Cmd) {
	input := strings.TrimSpace(m.input)
	if input == "" {
		return m, nil
	}
	
	// Adiciona comando ao output
	m.output = append(m.output, "$ "+input)
	
	// Parse dos argumentos
	args := strings.Fields(input)
	
	// Captura a saída do Cobra
	var buf bytes.Buffer
	m.rootCmd.SetOut(&buf)
	m.rootCmd.SetErr(&buf)
	
	// Reseta args e executa comando Cobra
	m.rootCmd.SetArgs(args)
	
	// Executa o comando
	if err := m.rootCmd.Execute(); err != nil {
		m.output = append(m.output, "Error: "+err.Error())
	} else {
		// Pega a saída do comando
		output := strings.TrimSpace(buf.String())
		
		// Verifica comandos especiais
		if output == "__EXIT__" {
			m.output = append(m.output, "Goodbye! 👋")
			m.input = ""
			m.cursorPos = 0
			return m, tea.Quit
		}
		
		if output == "__CLEAR_SCREEN__" {
			m.output = []string{}
			m.input = ""
			m.cursorPos = 0
			return m, nil
		}
		
		// Adiciona saída normal ao output
		if output != "" {
			lines := strings.Split(output, "\n")
			m.output = append(m.output, lines...)
		}
	}
	
	// Limpa o input
	m.input = ""
	m.cursorPos = 0
	
	return m, nil
}