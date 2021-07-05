package post

import (
	"github.com/z-t-y/flogo/cli/auth"
	"testing"
)

func TestListPosts(t *testing.T) {
	t.Parallel()
	token, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	posts, err := GetPosts(token)
	if err != nil {
		t.Error(err)
	}
	if len(posts) < 1 {
		t.Error("expected to get posts, actual 0")
	}
}