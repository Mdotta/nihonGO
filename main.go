package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
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

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
		m.keymap.quit,
	})
}

func getNewKana(pool map[string]string) (string, []string) {
	kana, romaji := getRandomKana(pool)
	except := []string{kana}
	alts := []string{romaji}

	n := 2
	for i := 0; i < n; i++ {
		k, r := getRandomKana(pool, except...)
		alts = append(alts, r)
		except = append(except, k)
	}

	rand.Shuffle(len(alts), func(i, j int) { alts[i], alts[j] = alts[j], alts[i] })
	return kana, alts
}

type keymap struct {
	up     key.Binding
	down   key.Binding
	choose key.Binding
	quit   key.Binding
}

func getHiraganaMap() map[string]string {
	return map[string]string{
		"あ": "a",
		"い": "i",
		"う": "u",
		"え": "e",
		"お": "o",
		"か": "ka",
		"き": "ki",
		"く": "ku",
		"け": "ke",
		"こ": "ko",
		"さ": "sa",
		"し": "shi",
		"す": "su",
		"せ": "se",
		"そ": "so",
		"た": "ta",
		"ち": "chi",
		"つ": "tsu",
		"て": "te",
		"と": "to",
	}
}

func getRandomKana(pool map[string]string, except ...string) (string, string) {
	exceptMap := make(map[string]bool)
	for _, k := range except {
		exceptMap[k] = true
	}

	for k := range pool {
		if !exceptMap[k] {
			return k, pool[k]
		}
	}
	return "", ""
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
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

func main() {
	model := model{
		pool: getHiraganaMap(),
		keymap: keymap{
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
		},
		help: help.New(),
	}
	model.current_kana, model.current_alternatives = getNewKana(model.pool)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
