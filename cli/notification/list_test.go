package notification

import (
	"testing"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestNotificationList(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	_, err = GetNotificationList(accessToken)
	if err != nil {
		t.Error(err)
	}
}
