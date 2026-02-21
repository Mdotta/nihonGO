package game

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	up       key.Binding
	down     key.Binding
	choose   key.Binding
	previous key.Binding
}

func (m flashcardGameModel) KeyList() []key.Binding {
	return []key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
		m.keymap.previous,
	}
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
		previous: key.NewBinding(
			key.WithKeys(tea.KeyBackspace.String()),
			key.WithHelp("Backspace", "go back"),
		),
	}
}
