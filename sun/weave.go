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

	for _, dir := range entries {
		if dir.Type().IsDir() && IsHack(dir.Name()) {
			return dir, nil
		}
	}

	return nil, &ErrNoHackDir{project: w.Project}
}

func (w *Weave) Hacks() ([]*Hack, error) {
	if w.hacks != nil {
		return w.hacks, nil
	}

	hackDir, err := w.GetHackDir()
	if err != nil {
		panic(err)
	}

	dirFullPath, err := filepath.Abs(hackDir.Name())
	if err != nil {
		panic(err)
	}

	var hacks []*Hack

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
			hacks = append(hacks, h)
		}
		return nil
	})

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

// Returns all files and directories under a Project.
func (w *Weave) Files() ([]os.DirEntry, error) {
	projectRoot := fmt.Sprintf("%s/%s", WeavesHome(), w.Project)
	return os.ReadDir(projectRoot)
}
