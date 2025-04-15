package sun

import (
	"fmt"
	"os"
)

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
