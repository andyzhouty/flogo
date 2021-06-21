package cmd

import (
	"os/exec"
	"strings"
	"testing"
)

func testArgs(args ...string) *exec.Cmd {
	argsCopy := args
	for _, arg := range argsCopy {
		args = append([]string{"run", mainGo}, arg)
	}
	return exec.Command("go", args...)
}

func TestRoot(t *testing.T) {
	t.Parallel()
	out, err := testArgs("").Output()
	if err != nil {
		t.Errorf("%s: %s", t.Name(), out)
	}
	if !strings.Contains(string(out), rootCmd.Long) {
		t.Errorf("%s: not containing the description", t.Name())
	}
}
