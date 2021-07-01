package cmd

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Run("no arg provided", func(t *testing.T) {
		t.Parallel()
		cmd := testArgs("config", "get")
		_, err := cmd.Output()
		if err == nil {
			t.Error("expected: exit status 1, actual: <nil>")
		}
	})
	t.Run("get flog url", func(t *testing.T) {
		testArgs("config", "set", "flog_url", "http://flog-web.herokuapp.com").Run()
		cmd := testArgs("config", "get", "flog_url")
		output, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		if len(output) == 0 {
			t.Error("error no output")
		}
		output = output[:len(output)-1] // avoid output errors
		if string(output) != "http://flog-web.herokuapp.com" {
			t.Errorf("expected %s, actual %s", "http://flog-web.herokuapp.com", output)
		}
	})
}
