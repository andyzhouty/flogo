package config

import (
	"reflect"
	"testing"

	"github.com/z-t-y/flogo/utils"
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
		command := utils.SetArgs("config", "set", "flog_url", "http://flog-web.herokuapp.com")
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
		if !reflect.DeepEqual(output, []byte("http://flog-web.herokuapp.com")) {
			t.Errorf("expected %s, actual %s", "http://flog-web.herokuapp.com", output)
		}
	})
}
