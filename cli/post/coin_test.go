package post

import (
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
)

func TestCoin(t *testing.T) {
	// generate a post
	token, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flogo Post Unit Test - " + strconv.FormatInt(time.Now().UnixNano(), 10)
	content := title
	post, err := UploadPost(title, content, token)
	if err != nil {
		t.Error(err)
	}
	// test if the GiveCoin function is correctly functioning
	post, err = GiveCoin(post.ID, 1, token)
	if err != nil {
		t.Error(err)
	}
	if post.Coins != 1 {
		t.Errorf("TestCoin: expect post.Coins to be 1, actual %d", post.Coins)
	}
}
