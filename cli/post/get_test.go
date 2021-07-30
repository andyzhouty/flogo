package post

import (
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestGetPost(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flogo Post Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := UploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	post2, err := GetPost(accessToken, post.ID)
	if err != nil {
		t.Error(err)
	}
	if post.ID != post2.ID || post.Title != post2.Title {
		t.Error("TestGetPost: expect post to equal post2")
	}
}
