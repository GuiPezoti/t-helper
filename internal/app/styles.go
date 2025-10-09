package app

import "github.com/charmbracelet/lipgloss"

// Paleta de cores base
var (
	ColorPrimary   = lipgloss.Color("#FF00FF") // Magenta
	ColorSecondary = lipgloss.Color("#AAAAAA") // Cinza
	ColorTitle     = lipgloss.Color("#FFFFFF") // Branco puro
)

// Estilos globais reutilizáveis
var (
	CommandStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary)

	DescriptionStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	TitleStyle = lipgloss.NewStyle().
			Foreground(ColorTitle).
			Bold(true).
			Underline(true)

	// Exemplo: usado em caixas ou headers
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 2)
)
