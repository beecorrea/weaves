package sun

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// A Weave extends a Moon Project with additional features.
type Weave struct {
	Project string  // Name of the Moon project for this Weave
	hacks   []*Hack // Hacks contained in this Weave
}

func (w *Weave) GetHackDir() (os.DirEntry, error) {
	entries, err := w.Files()
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.Type().IsDir() && IsHack(e.Name()) {
			return e, nil
		}
	}

	return nil, &ErrNoHackDir{project: w.Project}
}

// Returns all Hacks for a given Weave.
//
// If this the Hacks are loaded in memory by a previous call to this function,
// it'll return that value. Otherwise, it calls filepath.Walkdir using the
// Weave's hack directory as the root.
func (w *Weave) Hacks() ([]*Hack, error) {
	if w.hacks != nil {
		return w.hacks, nil
	}

	hackDir, err := w.GetHackDir()
	if err != nil {
		panic(err)
	}

	var hacks []*Hack
	dirFullPath := fmt.Sprintf("%s/%s", w.Root(), hackDir.Name())
	filepath.WalkDir(dirFullPath, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		info, err := entry.Info()
		if err != nil {
			panic(err)
		}

		if info.Mode().IsRegular() {
			h := &Hack{Name: entry.Name(), Path: path}
			h.runtime = h.Runtime()
			hacks = append(hacks, h)
		}
		return nil
	})

	w.hacks = hacks
	return hacks, nil
}

// Returns the value in WEAVES_HOME env var or panics.
func WeavesHome() string {
	const defaultWeavesRoot string = "WEAVES_HOME"
	home, exists := os.LookupEnv(defaultWeavesRoot)
	if !exists {
		msg := fmt.Sprintf("env %s not defined, exiting", defaultWeavesRoot)
		panic(msg)
	}
	return home
}

// Returns the root of the Weave
func (w *Weave) Root() string {
	return fmt.Sprintf("%s/%s", WeavesHome(), w.Project)
}

// Returns all files and directories under a Weave.
func (w *Weave) Files() ([]os.DirEntry, error) {
	projectRoot := fmt.Sprintf("%s/%s", WeavesHome(), w.Project)
	return os.ReadDir(projectRoot)
}
