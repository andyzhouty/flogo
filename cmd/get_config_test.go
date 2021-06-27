package cmd

import (
	"github.com/z-t-y/flogo/utils"
	"reflect"
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
		cmd := testArgs("config", "get", "flog_url")
		output, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		output = output[:len(output)-1] // avoid potential errors
		if !reflect.DeepEqual(output, []byte(utils.DefaultConfig.FlogURL)) {
			t.Errorf("expected %s, actual %s", utils.DefaultConfig.FlogURL, output)
		}
	})
}
