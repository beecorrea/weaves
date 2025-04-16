package main

import (
	"log"
	"os"

	"github.com/beecorrea/weaves/darlene"
	"github.com/beecorrea/weaves/miranda/pkg/hackrun"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(hackrun.InitModel("darlene"))
	m, err := p.Run()
	if err != nil {
		panic(err)
	}

	hackRun := m.(hackrun.HackRun)
	selected := hackRun.Selected
	if selected == nil {
		log.Println("no hack selected")
		os.Exit(0)
	}
	hr := darlene.Prepare(selected).WithOutput()
	if err := hr.Run(); err != nil {
		panic(err)
	}
}
