package self

import (
	"os"
	"testing"

	"github.com/z-t-y/flogo/cli/auth"
)

var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func TestGetSelf(t *testing.T) {
	token, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}

	self, err := GetSelf(token)
	if err != nil {
		t.Error(err)
	}
	if self.Name != username {
		t.Errorf("TestGetSelf: expected username to be %s, got %s", username, self.Name)
	}
}
