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
	quit   key.Binding
}

func mapKeys() keymap {
	return keymap{
		up: key.NewBinding(
			key.WithKeys(tea.KeyUp.String()),
			key.WithHelp(tea.KeyUp.String(), "go up"),
		),
		down: key.NewBinding(
			key.WithKeys(tea.KeyDown.String()),
			key.WithHelp(tea.KeyDown.String(), "go down"),
		),
		choose: key.NewBinding(
			key.WithKeys(tea.KeySpace.String()),
			key.WithHelp("Space", "choose"),
		),
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

type MenuModel struct {
	cursor           int
	keymap           keymap
	help             help.Model
	gamemodes        []string
	selectedGamemode string
	quitting         bool
}

func (m MenuModel) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
		m.keymap.quit,
	})
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true

		case key.Matches(msg, m.keymap.choose):
			m.selectedGamemode = m.gamemodes[m.cursor]

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
	s := "NihonGo"
	s += "\nGamemode selection:"
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
	keymap := mapKeys()
	model.keymap = model.keymap

	return model
}
