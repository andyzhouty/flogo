package config

import (
	"github.com/z-t-y/flogo/utils"
	"reflect"
	"testing"
)

func TestSetConfig(t *testing.T) {
	t.Run("no arg provided", func(t *testing.T) {
		command := utils.SetArgs("config", "set")
		_, err := command.Output()
		if err == nil {
			t.Error("expected: exit status 1, actual: <nil>")
		}
	})
	t.Run("1 arg provided", func(t *testing.T) {
		command := utils.SetArgs("config", "set", "flog_url")
		_, err := command.Output()
		if err == nil {
			t.Error("expected: exit status 1, actual: <nil>")
		}
	})
	t.Run("set flog url", func(t *testing.T) {
		command := utils.SetArgs("config", "set", "flog_url", "http://localhost:5000")
		_, err := command.Output()
		if err != nil {
			t.Error(err)
		}
		command = utils.SetArgs("config", "get", "flog_url")
		output, err := command.Output()
		if err != nil {
			t.Error(err)
		}
		output = output[:len(output)-1]
		if !reflect.DeepEqual(output, []byte("http://localhost:5000")) {
			t.Errorf("expected %s, actual %s", "http://localhost:5000", output)
		}
	})
}
