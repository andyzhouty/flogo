package cmd

import (
	"strconv"
	"testing"
	"time"
)

func TestPublish(t *testing.T) {
	token, err := getAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flog Unittest - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := uploadPost(title, content, token)
	if err != nil {
		t.Error(err)
	}
	if post.Title != title {
		t.Errorf("TestPublish: expected title %s, actual %s", title, post.Title)
	}
}
