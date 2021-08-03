package post

import (
	"testing"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestSelfPosts(t *testing.T) {
	t.Parallel()
	token, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	posts, err := GetPostsFrom("/api/v3/self/posts", token)
	if err != nil {
		t.Error(err)
	}
	if len(posts) < 1 {
		t.Error("expected to get posts, actual 0")
	}
}

func TestAllPosts(t *testing.T) {
	t.Parallel()
	posts, err := GetPostsFrom("/api/v3/post/all", "")
	if err != nil {
		t.Error(err)
	}
	for _, post := range posts {
		if post.Private {
			t.Error("expected all returned posts to be public")
		}
	}
}
