package cmd

import (
	"strconv"
	"testing"
	"time"
)

func TestGetPost(t *testing.T) {
	t.Parallel()
	accessToken, err := getAccessToken(username, password)
	title := "Flog Unittest - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := uploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	post2, err := getPost(accessToken, post.ID)
	if err != nil {
		t.Error(err)
	}
	if post.ID != post2.ID || post.Title != post2.Title {
		t.Errorf("TestGetPost: expected post to equal post2")
	}
}
