package config

import (
	"testing"

	"github.com/z-t-y/flogo/utils"
)

func TestGetConfig(t *testing.T) {
	t.Run("no arg provided", func(t *testing.T) {
		command := utils.SetArgs("config", "get")
		_, err := command.Output()
		if err == nil {
			t.Error("expected: exit status 1, actual: <nil>")
		}
	})
	t.Run("get flog url", func(t *testing.T) {
		utils.SetArgs("config", "set", "flog_url", "http://flog-web.herokuapp.com").Run()
		command := utils.SetArgs("config", "get", "flog_url")
		output, err := command.Output()
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
