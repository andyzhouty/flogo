package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Parallel()
	expected := fmt.Sprintf("Go: %s\nFlogo: %s\n", runtime.Version(), version)
	out, err := exec.Command("go", "run", "../main.go", "version").Output()
	if err != nil {
		t.Errorf("TestVersion: %s", err.Error())
	}
	if string(out) != expected {
		t.Errorf("TestVersion: expected %s, actual %s", expected, out)
	}
}
