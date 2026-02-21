package main

import (
	"fmt"
	"hiragana-guesser/modelStack"
	"hiragana-guesser/screens/flashcard/menu"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type parentModel struct {
	stack        modelStack.ModelStack
	globalKeymap GlobalKeymap
	help         help.Model
}

func (m parentModel) Init() tea.Cmd {
	return m.stack.Init()
}

func (m parentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.globalKeymap.quit):
			return m, tea.Quit
		}
	}
	updated, cmd := m.stack.Update(msg)
	m.stack = updated.(modelStack.ModelStack)
	return m, cmd
}

func (m parentModel) View() string {
	view := m.stack.View()
	view += "\n" + m.help.ShortHelpView([]key.Binding{
		m.globalKeymap.quit,
	})
	return view
}

func main() {
	baseModel := menu.NewModel()
	stack := modelStack.New(baseModel)
	globalKeys := mapGlobalKeys()

	parent := parentModel{
		stack:        stack,
		globalKeymap: globalKeys,
		help:         help.New(),
	}

	p := tea.NewProgram(parent)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
