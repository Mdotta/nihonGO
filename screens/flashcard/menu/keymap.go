package menu

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keymap struct {
	up       key.Binding
	down     key.Binding
	choose   key.Binding
	options  key.Binding
	previous key.Binding
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
		options: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "toggle options"),
		),
		previous: key.NewBinding(
			key.WithKeys(tea.KeyBackspace.String()),
			key.WithHelp("Backspace", "go back"),
		),
	}
}

func (m MenuModel) KeyList() []key.Binding {
	return []key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
		m.keymap.options,
		m.keymap.previous,
	}
}
