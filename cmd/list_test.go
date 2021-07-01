package cmd

import "testing"

func TestListPosts(t *testing.T) {
	t.Parallel()
	token, err := getAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	posts, err := getPosts(token)
	if err != nil {
		t.Error(err)
	}
	if len(posts) < 1 {
		t.Error("expected to get posts, actual 0")
	}
}