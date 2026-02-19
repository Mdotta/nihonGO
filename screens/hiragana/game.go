package hiragana

import (
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// timer    timer.Model
	score                int
	cursor               int
	pool                 map[string]string
	current_kana         string
	current_alternatives []string
	keymap               keymap
	help                 help.Model
	quitting             bool
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
		m.keymap.quit,
	})
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

type keymap struct {
	up     key.Binding
	down   key.Binding
	choose key.Binding
	quit   key.Binding
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit

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
		}

	}

	return m, nil
}

func (m Model) View() string {
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

func NewModel() Model {
	model := Model{
		pool:   getHiraganaMap(),
		keymap: mapKeys(),
		help:   help.New(),
	}
	model.current_kana, model.current_alternatives = getNewKana(model.pool)

	return model
}
