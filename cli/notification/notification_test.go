package notification

import (
	"os"
	"testing"

	"github.com/z-t-y/flogo/utils"
)

var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func TestNotificationCommand(t *testing.T) {
	cmd := utils.SetArgs("notification")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
}
