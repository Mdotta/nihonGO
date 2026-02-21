package menu

func (m MenuModel) helpView() string {
	return "\n" + m.help.ShortHelpView(m.KeyList())
}

type resetCursorMsg struct {
}
