package hackrun

import (
	"log"
	"os"

	"github.com/beecorrea/weaves/darlene"
	tea "github.com/charmbracelet/bubbletea"
)

func Wear(weave string) {
	p := tea.NewProgram(InitModel(weave))
	m, err := p.Run()
	if err != nil {
		panic(err)
	}

	hackRun := m.(HackRun)
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
