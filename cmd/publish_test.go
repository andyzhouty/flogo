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
	statusCode := uploadPost(title, content, token)
	if statusCode != 200 {
		t.Errorf("TestPublish: expected status code 200, actual %d", statusCode)
	}
}
