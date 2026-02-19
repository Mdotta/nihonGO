package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type gameFunc func() tea.Model

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("q", "quit"),
)

func (m appModel) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		quitKeys,
	})
}
