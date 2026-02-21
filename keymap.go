package main

import "github.com/charmbracelet/bubbles/key"

type GlobalKeymap struct {
	quit key.Binding
}

func mapGlobalKeys() GlobalKeymap {
	return GlobalKeymap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

func (gk GlobalKeymap) KeyList() []key.Binding {
	return []key.Binding{
		gk.quit,
	}
}
