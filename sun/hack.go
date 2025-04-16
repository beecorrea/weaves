package sun

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	defaultRuntime string = "/bin/sh"
)

// A Hack is a script that automates some task at the Project level.
type Hack struct {
	Name    string // Name of the Hack script
	Path    string // Path to the Hack script
	runtime string // Interpreter used to run the Hack script, usually /bin/sh.
}

// Weave doesn't contain a Hack directory
type ErrNoHackDir struct {
	project string
}

func (e *ErrNoHackDir) Error() string {
	return fmt.Sprintf("%s doesnt have a hack dir", e.project)
}

// Returns true if a dirname is "hack", false otherwise.
func IsHack(dirName string) bool {
	return dirName == "hack"
}

func (h *Hack) Runtime() string {
	if h.runtime != "" {
		return h.runtime
	}

	file, err := os.Open(h.Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	hasLine := scanner.Scan()
	if !hasLine {
		return defaultRuntime
	}
	shebang := scanner.Text()
	runtime := strings.ReplaceAll(shebang, "#!", "")
	return runtime
}
