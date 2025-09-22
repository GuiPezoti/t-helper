package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WelcomeFinishedMsg struct{}

type WelcomeModel struct {
	width 	int
	height 	int
}

func NewWelcome() WelcomeModel {
	return WelcomeModel{}
}

func (wm WelcomeModel) Init() tea.Cmd {
	return nil
}

func (wm WelcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		wm.width = msg.Width
		wm.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return wm, func() tea.Msg { return WelcomeFinishedMsg{} }
		}
	}
	return wm, nil
}

func  (wm WelcomeModel) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Padding(1, 0)
	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#AAA"))

	title := titleStyle.Render("Bem-vindo ao T-Helper!")
	subtitle := subtitleStyle.Render("Sua ferramenta de CLI para automação.\n\nPressione 'Enter' para continuar...")
	return lipgloss.Place(
		wm.width, wm.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Left, title, subtitle),
	)
}