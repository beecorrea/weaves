package main

import (
	"os"

	"github.com/beecorrea/weaves/miranda/cmd"
)

func main() {
	err := cmd.MirandaCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
