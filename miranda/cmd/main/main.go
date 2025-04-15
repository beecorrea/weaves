package main

import (
	"github.com/beecorrea/weaves/miranda/pkg/hackerrun.go"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	panic(err)
	// }

	project := "miranda"
	p := tea.NewProgram(hackerrun.InitModel(project))
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
