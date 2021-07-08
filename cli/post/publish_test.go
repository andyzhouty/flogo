package post

import (
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestPublish(t *testing.T) {
	token, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flogo Post Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := UploadPost(title, content, token)
	if err != nil {
		t.Error(err)
	}
	if post.Title != title {
		t.Errorf("TestPublish: expected title %s, actual %s", title, post.Title)
	}
}
