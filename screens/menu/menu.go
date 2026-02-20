package menu

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	up     key.Binding
	down   key.Binding
	choose key.Binding
}

type MenuModel struct {
	cursor           int
	keymap           keymap
	help             help.Model
	gamemodes        []string
	SelectedGamemode string
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.choose):
			m.SelectedGamemode = m.gamemodes[m.cursor]

		case key.Matches(msg, m.keymap.up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keymap.down):
			if m.cursor < len(m.gamemodes)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m MenuModel) View() string {
	s := "Gamemode selection:"
	for i, mode := range m.gamemodes {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += "\n" + cursor + mode
	}

	s += m.helpView()
	return s
}

func NewModel(availableGamemodes []string) MenuModel {
	model := MenuModel{
		cursor:    0,
		gamemodes: availableGamemodes,
		keymap:    mapKeys(),
		help:      help.New(),
	}

	return model
}
