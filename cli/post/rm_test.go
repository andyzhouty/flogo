package post

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestRmPost(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flog Post Unit Test - " + strconv.FormatInt(time.Now().UnixNano(), 10)
	content := title
	post, err := UploadPost(title, content, accessToken)
	if err != nil {
		t.Error(err)
	}
	err = RmPost(accessToken, post.ID)
	if err != nil {
		t.Error(err)
	}
	p, err := GetPost(accessToken, post.ID)
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Errorf("TestRmPost: expected to get a 404 response, actual %v", p)
	}
}
