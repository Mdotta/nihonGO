package main

import (
	"fmt"
	"hiragana-guesser/screens/menu"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type screen int

const (
	menuScreen screen = iota
	gameScreen
)

type appModel struct {
	currentScreen screen
	menuModel     menu.MenuModel
	gameModel     tea.Model
}

func (m appModel) Init() tea.Cmd {
	//return menu init
	return m.menuModel.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	quitting := false
	switch m.currentScreen {
	case menuScreen:
		updated, cmd := m.menuModel.Update(msg)
		m.menuModel = updated.(menu.MenuModel)
		//TODO: check if gamemode was selected
		return m, cmd
	case gameScreen:
		updated, cmd := m.gameModel.Update(msg)
		m.gameModel = updated

		return m, cmd
	}

	return m, nil
}

func (m appModel) View() string {
	switch m.currentScreen {
	case menuScreen:
		return m.menuModel.View()
	case gameScreen:
		return m.gameModel.View()
	}
	return ""
}

var gamemodes = []string{
	"Hiragana",
	"Katakana",
	"Kanji",
}

func main() {
	model := appModel{
		currentScreen: menuScreen,
		menuModel:     menu.NewModel(gamemodes),
		gameModel:     nil,
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
