package game

import (
	"hiragana-guesser/modelStack"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type flashcardGameModel struct {
	// timer    timer.Model
	score                int
	cursor               int
	pool                 map[string]string
	current_kana         string
	current_alternatives []string
	keymap               keymap
	help                 help.Model
}

func (m flashcardGameModel) Init() tea.Cmd {
	return nil
}

func (m flashcardGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keymap.up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keymap.down):
			if m.cursor < len(m.current_alternatives)-1 {
				m.cursor++
			}

		case key.Matches(msg, m.keymap.choose):
			if m.pool[m.current_kana] == m.current_alternatives[m.cursor] {
				m.score++
			}
			m.cursor = 0
			m.current_kana, m.current_alternatives = getNewKana(m.pool)
		case key.Matches(msg, m.keymap.previous):
			return m, modelStack.Pop()
		}

	}

	return m, nil
}

func (m flashcardGameModel) View() string {
	s := "score: " + strconv.Itoa(m.score)
	s += "\n" + m.current_kana
	for i, alt := range m.current_alternatives {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += "\n"
		s += cursor
		s += alt
	}
	s += m.helpView()
	return s
}

func NewModel(pool map[string]string, options ...Option) flashcardGameModel {
	model := flashcardGameModel{
		pool:   pool,
		keymap: mapKeys(),
		help:   help.New(),
	}

	for _, option := range options {
		option(&model)
	}

	model.current_kana, model.current_alternatives = getNewKana(model.pool)

	return model
}
