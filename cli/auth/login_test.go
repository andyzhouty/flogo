package auth

import (
	"fmt"
	"os"
	"testing"
)

var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func TestLogin(t *testing.T) {
	fmt.Println(username, password)
	t.Parallel()
	t.Run("using Username and Password", func(t *testing.T) {
		t.Parallel()
		err := useUsernamePassword(username, password)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("test getting token", func(t *testing.T) {
		t.Parallel()
		_, err := GetAccessToken(username, password)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("test logging in using token", func(t *testing.T) {
		t.Parallel()
		token, _ := GetAccessToken(username, password)
		returnedUsername, err := verifyToken(token)
		if err != nil {
			t.Error(err)
		} else if returnedUsername != username {
			t.Errorf("%s: returned Username does not match the actual one", t.Name())
		}
	})
}
