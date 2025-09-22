package models

import ( 
	tea "github.com/charmbracelet/bubbletea"
)

type AppState int

const (
	Welcome  AppState = iota
    Finished
)

type AppModel struct {
	state AppState
	// gitInstalled	bool
	// directory		map[string]string
	// gitToken		string
	width          	int
    height         	int
	views 			map[AppState]tea.Model
	error			error
}

func NewAppModel() AppModel {
	return AppModel{
		state: Welcome,
		views: make(map[AppState]tea.Model),
	}
}

func (app AppModel) Init() tea.Cmd{
	return nil
}

func (app AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		app.width = msg.Width
		app.height = msg.Height

	case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
			return app, tea.Quit
		}

	case WelcomeFinishedMsg:
		app.state = Finished
		return app, nil
	}

	if _, ok := app.views[app.state]; !ok {
		switch app.state {
			//Adicionar novas views como case switch abaixo
		case Welcome:
			v:= NewWelcome()
			app.views[Welcome] = v
			return app, v.Init()
		}
	}
	var cmd tea.Cmd
	app.views[app.state], cmd = app.views[app.state].Update(msg)
	return app, cmd	
}

func (app AppModel) View() string {
	if v, ok := app.views[app.state]; ok {
		switch m := v.(type) {
		case WelcomeModel:
			return m.View()
		}
	}
	return "Carregando..."
}
