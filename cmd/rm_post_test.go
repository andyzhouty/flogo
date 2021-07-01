package cmd

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRmPost(t *testing.T) {
	t.Parallel()
	accessToken, err := getAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flog Unittest - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := uploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	err = rmPost(accessToken, post.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = getPost(accessToken, post.ID)
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Error("TestRmPost: expected to get a 404 response.")
	}
}
