package hackrun

import (
	"github.com/beecorrea/weaves/sun"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model for a HackRun instance
type HackRun struct {
	List     list.Model
	Cursor   int
	Selected *sun.Hack
}

func InitModel(project string) HackRun {
	w := &sun.Weave{Project: project}
	hacks, err := w.Hacks()
	if err != nil {
		panic(err)
	}
	l := NewHackList(hacks)

	hr := HackRun{
		List:     l,
		Cursor:   0,
		Selected: nil,
	}

	return hr
}

func (hr HackRun) Init() tea.Cmd {
	return nil
}

func (hr HackRun) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return hr, tea.Quit
		// Choose a Hack and quit to main
		case key.Matches(msg, DefaultKeyMap.Select):
			return hr.Select()
		}
	}

	var cmd tea.Cmd
	hr.List, cmd = hr.List.Update(msg)
	return hr, cmd
}

func (hr HackRun) View() string {
	return hr.List.View()
}

func (hr HackRun) Select() (HackRun, tea.Cmd) {
	item := hr.List.SelectedItem().(Item)
	h := sun.Hack(item)
	hr.Selected = &h
	return hr, tea.Quit
}
