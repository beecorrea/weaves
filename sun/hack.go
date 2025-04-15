package sun

import "fmt"

// A Hack is a script that automates some task at the Project level.
type Hack struct {
	Name string // Name of the Hack script
	Path string // Path to the Hack script
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
