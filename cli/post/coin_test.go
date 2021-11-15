package post

import (
	"net/http"
	"net/url"
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

	data := url.Values{}
	data.Add("username", "abc")
	data.Add("password", "abc")
	data.Add("email", "abc@example.com")
	http.PostForm("http://localhost:5000/api/v3/register",  data)
	// test if the GiveCoin function is correctly functioning
	token, err = auth.GetAccessToken("abc", "abc")
	if err != nil {
		t.Error(err)
	}
	post, err = GiveCoin(post.ID, 1, token)
	if err != nil {
		t.Error(err)
	}
	if post.Coins != 1 {
		t.Errorf("TestCoin: expect post.Coins to be 1, actual %d", post.Coins)
	}
}
