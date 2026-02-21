package modelStack

import (
	"hiragana-guesser/modelStack/internal/stack"

	tea "github.com/charmbracelet/bubbletea"
)

type ModelStack struct {
	current    tea.Model
	lastResize tea.WindowSizeMsg
	stack      stack.Stack[tea.Model]
}

type PushModel struct {
	Model tea.Model
}

type PopModel struct {
	Msgs []tea.Msg
}

func New(initialModel tea.Model) ModelStack {
	return ModelStack{
		current: initialModel,
		stack:   stack.Stack[tea.Model]{},
	}
}

func Push(m tea.Model) tea.Cmd {
	return func() tea.Msg {
		return PushModel{
			Model: m,
		}
	}
}

func Pop(msgs ...tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return PopModel{
			Msgs: msgs,
		}
	}
}

func (ms *ModelStack) updateCurrent(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	ms.current, cmd = ms.current.Update(msg)
	return cmd
}

//tea funcs

func (ms ModelStack) Init() tea.Cmd {
	return ms.current.Init()
}

func (ms ModelStack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ms.lastResize = msg
	case PushModel:
		ms.stack.Push(ms.current)
		ms.current = msg.Model
		cmd := ms.current.Init()
		cmd2 := ms.updateCurrent(ms.lastResize)
		return ms, tea.Batch(cmd, cmd2)
	case PopModel:
		if ms.stack.Peek() == nil {
			return ms, nil
		}
		ms.current = ms.stack.Pop().Value
		cmds := make(tea.BatchMsg, 2+len(msg.Msgs))
		cmds[0] = ms.current.Init()
		cmds[1] = ms.updateCurrent(ms.lastResize)
		for i, msg := range msg.Msgs {
			cmds[i+2] = ms.updateCurrent(msg)
		}
		return ms, tea.Batch(cmds...)
	}

	cmd := ms.updateCurrent(msg)
	return ms, cmd
}

func (ms ModelStack) View() string {
	return ms.current.View()
}
