package cmd

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

var mainGo = "../main.go"
var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func testArgs(args ...string) *exec.Cmd {
	argsCopy := args
	args = []string{"run", mainGo}
	for _, arg := range argsCopy {
		args = append(args, arg)
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
