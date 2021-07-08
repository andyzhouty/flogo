package comment

import (
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
	"github.com/z-t-y/flogo/cli/post"
)

func TestRmComment(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flog Post Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	content := title
	post, err := post.UploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	content = "Flog Comment Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	comment, err := addComment(accessToken, content, post.ID, 0)
	if err != nil {
		t.Error(err)
	}
	err = rmComment(accessToken, comment.ID)
	if err != nil {
		t.Error(err)
	}
}
