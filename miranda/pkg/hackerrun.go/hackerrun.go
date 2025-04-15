package hackerrun

import (
	"fmt"

	"github.com/beecorrea/weaves/sun"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Model for a Hackerrun instance
type Hackerrun struct {
	Weaves   []*sun.Weave
	Cursor   int
	Selected map[int]*sun.Weave
}

func InitModel(project string) Hackerrun {
	w := sun.Weave{Project: project}

	hr := Hackerrun{Weaves: []*sun.Weave{&w}, Selected: make(map[int]*sun.Weave), Cursor: 1}
	// hacks, _ := w.Hacks()
	// for _, h := range hacks {
	// 	fmt.Println(h.Name, h.Path)
	// }
	return hr
}

func (hr Hackerrun) Init() tea.Cmd {
	return nil
}

func (hr Hackerrun) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// Move cursor up
		case key.Matches(msg, DefaultKeyMap.Up):
			if hr.Cursor > 0 {
				hr.Cursor--
			}

		// Move cursor down
		case key.Matches(msg, DefaultKeyMap.Down):
			if hr.Cursor < len(hr.Weaves) {
				hr.Cursor++
			}

		// Choose a Hack
		case key.Matches(msg, DefaultKeyMap.Select):
			_, ok := hr.Selected[hr.Cursor]
			if ok {
				delete(hr.Selected, hr.Cursor)
			} else {
				hr.Selected[hr.Cursor] = hr.Weaves[hr.Cursor]
			}

		// Quit app
		case key.Matches(msg, DefaultKeyMap.Quit):
			return hr, tea.Quit
		}
	}

	return hr, nil
}

func (hr Hackerrun) View() string {
	s := "HACKS\n\n"
	for _, w := range hr.Weaves {
		s += fmt.Sprintf("%s\n", w.Project)
		hacks, err := w.Hacks()
		if err != nil {
			panic(err)
		}

		for j, h := range hacks {
			cursor := " "
			if hr.Cursor == j {
				cursor = ">"
			}

			checked := " " // not selected
			if _, ok := hr.Selected[j]; ok {
				checked = "x" // selected!
			}

			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, h.Name)
		}
	}
	return s
}
