package darlene

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/beecorrea/weaves/sun"
)

// A HackRun associates a given Hack to its arguments and the exec.Cmd used to run it.
type HackRun struct {
	hack   *sun.Hack
	cmd    *exec.Cmd
	cmdOut bytes.Buffer
}

// Prepare creates a exec.Command that pipes to os.{Stdin, Stdout, Stderr}
// It's meant to be used along with hr.WithArgs and hr.Run.
func Prepare(h *sun.Hack) *HackRun {
	cmd := &exec.Cmd{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Path = h.Runtime()
	cmd.Args = []string{h.Runtime(), h.Path}

	return &HackRun{
		hack: h,
		cmd:  cmd,
	}
}

// WithArgs adds arguments to the HackRun Command.
//
// If there's no Command, WithArgs returns an error.
func (hr *HackRun) WithArgs(args ...string) *HackRun {
	hr.cmd.Args = append(hr.cmd.Args, args...)
	return hr
}

// WithOutput captures the command output into HackRun.cmdOut.
func (hr *HackRun) WithOutput() *HackRun {
	hr.cmd.Stdout = io.MultiWriter(hr.cmd.Stdout, &hr.cmdOut)
	return hr
}

// Run executes a HackRun and stores the result in the HackRun.cmdOut field.
func (hr *HackRun) Run() error {
	return hr.cmd.Run()
}

// Output returns hr.cmdOut whether it's empty or not.
func (hr *HackRun) Output() string {
	l := hr.cmdOut.Len()
	// Remove last byte to avoid printing whitespaces
	return string(hr.cmdOut.Bytes()[0 : l-1])
}
