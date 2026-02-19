package menu

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

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
	}
}

func (m MenuModel) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
	})
}

func (m MenuModel) KeyList() []key.Binding {
	return []key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
	}
}
