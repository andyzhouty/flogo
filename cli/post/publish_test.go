package post

import (
	"github.com/z-t-y/flogo/cli/auth"
	"strconv"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	t.Parallel()
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
