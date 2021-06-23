package cmd

import (
	"fmt"
	"runtime"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Parallel()
	expected := fmt.Sprintf("Go: %s\nFlog-CLI: %s\n", runtime.Version(), version)
	out, err := testArgs("version").Output()
	if err != nil {
		t.Errorf("TestVersion: %s", err.Error())
	}
	if string(out) != expected {
		t.Errorf("TestVersion: expected %s, actual %s", expected, out)
	}
}
