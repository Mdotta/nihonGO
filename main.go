package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type appModel struct {
	help help.Model
}

func (m appModel) Init() tea.Cmd {
	//return menu init
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m appModel) View() string {
	view := "NihonGo\n"
	view += "========\n\n"

	return view
}

func main() {
	model := appModel{
		help: help.New(),
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
