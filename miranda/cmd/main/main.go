package main

import (
	"log"
	"os"

	"github.com/beecorrea/weaves/miranda/pkg/hackerrun.go"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(hackerrun.InitModel("miranda"))
	hr, err := p.Run()
	if err != nil {
		panic(err)
	}

	selected, _ := hr.(hackerrun.Hackerrun)
	if selected.Selected == nil {
		log.Println("no hack selected")
		os.Exit(0)
	}
	log.Println("Running hack:", selected.Selected.Name)
}
