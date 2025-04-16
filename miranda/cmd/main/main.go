package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/beecorrea/weaves/miranda/pkg/hackerrun.go"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(hackerrun.InitModel("miranda"))
	m, err := p.Run()
	if err != nil {
		panic(err)
	}

	hackRun := m.(hackerrun.Hackerrun)
	selected := hackRun.Selected
	if selected == nil {
		log.Println("no hack selected")
		os.Exit(0)
	}
	// TODO(Bia): Move hackerrun to another package.
	log.Println("Running hack:", selected.Name)
	c := exec.Command("/bin/sh", selected.Path)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		panic(err)
	}
	log.Println("Finished")
}
