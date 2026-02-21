package menu

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	modelStack "hiragana-guesser/modelStack"
	"hiragana-guesser/screens/flashcard/game"
)

type gameOptions struct {
	name    string
	enabled bool
	option  game.Option
}

type gameMode struct {
	name string
	pool map[string]string
}

type MenuModel struct {
	cursor      int
	gamemodes   []gameMode
	showOptions bool
	options     []gameOptions

	keymap keymap
	help   help.Model
}

func (m MenuModel) Init() tea.Cmd {
	return func() tea.Msg {
		return resetCursorMsg{}
	}
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case resetCursorMsg:
		m.cursor = 0
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.choose):
			//TODO: this should be dynamic based on the gamemode selected
			switch m.showOptions {
			case true:
				option := m.options[m.cursor]
				option.enabled = !option.enabled
				m.options[m.cursor] = option
			case false:
				pool := m.gamemodes[m.cursor].pool
				options := []game.Option{}
				for _, option := range m.options {
					if option.enabled {
						options = append(options, option.option)
					}
				}
				gameModel := game.NewModel(pool, options...)
				return m, modelStack.Push(gameModel)
			}

		case key.Matches(msg, m.keymap.options):
			m.cursor = 0
			m.showOptions = !m.showOptions

		case key.Matches(msg, m.keymap.up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keymap.down):
			limit := len(m.gamemodes)
			if m.showOptions {
				limit = len(m.options)
			}
			if m.cursor < limit-1 {
				m.cursor++
			}

		case key.Matches(msg, m.keymap.previous):
			return m, modelStack.Pop()
		}
	}

	return m, nil
}

func (m MenuModel) View() string {
	s := ""
	switch m.showOptions {
	case true:
		s += "Options:"
		for i, option := range m.options {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			enabled := "[ ]"
			if option.enabled {
				enabled = "[X]"
			}
			s += "\n" + cursor + enabled + " " + option.name
		}
	case false:
		s += "Gamemode selection:"
		for i, mode := range m.gamemodes {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += "\n" + cursor + mode.name
		}
	}

	s += m.helpView()
	return s
}

func NewModel() MenuModel {
	model := MenuModel{
		cursor: 0,
		gamemodes: []gameMode{
			{"Hiragana", game.HiraganaPool},
			{"Katakana", game.KatakanaPool},
		},
		keymap: mapKeys(),
		help:   help.New(),
		options: []gameOptions{
			{"Reverse mode", false, game.WithReverseMode()},
			{"Use words", false, game.WithReverseMode()},
		},
	}

	return model
}
