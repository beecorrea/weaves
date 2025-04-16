package darlene

import (
	"strings"
	"testing"

	"github.com/beecorrea/weaves/sun"
)

const (
	TestProject string = "darlene"
)

func TestExecutesSuccessfully(t *testing.T) {
	w := &sun.Weave{Project: TestProject}
	hacks, err := w.Hacks()
	if err != nil {
		panic(err)
	}

	testCases := map[string]string{
		"new_script":    "This is a new script",
		"sample":        "This is a sample hack script",
		"sample2":       "This is another hack script",
		"python_script": "This is a Python script!",
	}

	// For each Hack, create a HackRun and check if the output matches the expected.
	for _, hack := range hacks {
		name := hack.Name
		hr := Prepare(hack).WithOutput()
		err := hr.Run()
		if err != nil {
			t.Errorf("error running '%s': %s", name, err.Error())
		}
		actual := strings.TrimSpace(hr.Output())

		if expected, ok := testCases[name]; !ok || expected != actual {
			t.Logf("output from '%s' is different from expected", name)
			t.Logf("\tactual: '%s'", actual)
			t.Logf("\texpected: '%s'", expected)
			t.FailNow()
		}
	}
}
