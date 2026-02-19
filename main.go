package main

import (
	"fmt"
	"hiragana-guesser/screens/hiragana"
	"hiragana-guesser/screens/menu"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	menuScreen screen = iota
	gameScreen
)

var gamemodes = []string{
	"Hiragana",
	"Katakana",
	"Kanji",
}

type appModel struct {
	currentScreen screen
	menuModel     menu.MenuModel
	gameModel     tea.Model
	help          help.Model
}

func (m appModel) Init() tea.Cmd {
	//return menu init
	return m.menuModel.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, quitKeys):
			if m.currentScreen == gameScreen {
				m.currentScreen = menuScreen
				m.gameModel = nil
				return m, nil
			}
			return m, tea.Quit
		}
	}

	switch m.currentScreen {
	case menuScreen:
		updated, cmd := m.menuModel.Update(msg)
		m.menuModel = updated.(menu.MenuModel)

		//TODO: check if gamemode was selected
		if m.menuModel.SelectedGamemode != "" {
			m.currentScreen = gameScreen
			m.gameModel = hiragana.NewModel()
			return m, m.gameModel.Init()
		}
		return m, cmd
	case gameScreen:

		updated, cmd := m.gameModel.Update(msg)
		m.gameModel = updated

		return m, cmd
	}

	return m, nil
}

func (m appModel) View() string {
	view := "NihonGo\n"
	switch m.currentScreen {
	case menuScreen:
		view += m.menuModel.View()
	case gameScreen:
		view += m.gameModel.View()
	}

	view += m.helpView()
	return view
}

func main() {
	model := appModel{
		currentScreen: menuScreen,
		menuModel:     menu.NewModel(gamemodes),
		gameModel:     nil,
		help:          help.New(),
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
