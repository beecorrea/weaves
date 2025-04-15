package hackerrun

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Quit   key.Binding
	Select key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/spacebar", "select"),
	),
}
