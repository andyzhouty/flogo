package post

import (
	"github.com/z-t-y/flogo/cli/auth"
	"strconv"
	"testing"
	"time"
)

func TestGetPost(t *testing.T) {
	t.Parallel()
	accessToken, err := auth.GetAccessToken(username, password)
	title := "Flog Post Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := UploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	post2, err := getPost(accessToken, post.ID)
	if err != nil {
		t.Error(err)
	}
	if post.ID != post2.ID || post.Title != post2.Title {
		t.Error("TestGetPost: expect post to equal post2")
	}
}
