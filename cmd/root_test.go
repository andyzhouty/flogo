package cmd

import (
	"os/exec"
	"strings"
	"testing"
)



func TestRoot(t *testing.T) {
	t.Parallel()
	out, err := exec.Command("go", "run", "../main.go").Output()
	if err != nil {
		t.Errorf("%s: %s", t.Name(), out)
	}
	if !strings.Contains(string(out), RootCmd.Long) {
		t.Errorf("%s: not containing the description", t.Name())
	}
}
