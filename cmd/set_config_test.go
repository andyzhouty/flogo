package cmd

import (
	"reflect"
	"testing"
)

func TestSetConfig(t *testing.T) {
	t.Run("no arg provided", func(t *testing.T) {
		t.Parallel()
		cmd := testArgs("config", "set")
		_, err := cmd.Output()
		if err == nil {
			t.Error("expected: exit status 1, actual: <nil>")
		}
	})
	t.Run("1 arg provided", func(t *testing.T) {
		t.Parallel()
		cmd := testArgs("config", "set", "flog_url")
		_, err := cmd.Output()
		if err == nil {
			t.Error("expected: exit status 1, actual: <nil>")
		}
	})
	t.Run("set flog url", func(t *testing.T) {
		cmd := testArgs("config", "set", "flog_url", "http://flog-web.herokuapp.com")
		_, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		cmd = testArgs("config", "get", "flog_url")
		output, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		output = output[:len(output)-1]
		if !reflect.DeepEqual(output, []byte("http://flog-web.herokuapp.com")) {
			t.Errorf("expected %s, actual %s", "http://flog-web.herokuapp.com", output)
		}
	})
}
