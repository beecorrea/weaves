package hackerrun

import (
	"github.com/beecorrea/weaves/sun"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model for a Hackerrun instance
type Hackerrun struct {
	Weaves   []*sun.Weave
	List     list.Model
	Cursor   int
	Selected *sun.Hack
}

func InitModel(project string) Hackerrun {
	w := &sun.Weave{Project: project}
	hacks, err := w.Hacks()
	if err != nil {
		panic(err)
	}
	l := NewHackList(hacks)

	hr := Hackerrun{
		Weaves:   []*sun.Weave{w},
		List:     l,
		Cursor:   0,
		Selected: nil,
	}

	return hr
}

func (hr Hackerrun) Init() tea.Cmd {
	return nil
}

func (hr Hackerrun) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return hr, tea.Quit
		// Choose a Hack and quit to main
		case key.Matches(msg, DefaultKeyMap.Select):
			item := hr.List.SelectedItem().(Item)
			h := sun.Hack(item)
			hr.Selected = &h
			return hr, tea.Quit
		}

	}

	var cmd tea.Cmd
	hr.List, cmd = hr.List.Update(msg)
	return hr, cmd
}

func (hr Hackerrun) View() string {
	return hr.List.View()
}
