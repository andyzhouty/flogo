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
	notifications, err := GetNotificationList(accessToken)
	if err != nil {
		t.Error(err)
	}
	if notifications[0].Message == "" {
		t.Error("TestNotificationList: empty message")
	}
}
