package cmd

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	username := os.Getenv("FLOG_USERNAME")
	password := os.Getenv("FLOG_PASSWORD")
	t.Run("using username and password", func(t *testing.T) {
		t.Parallel()
		err := useUsernamePassword(username, password)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("test getting token", func(t *testing.T) {
		t.Parallel()
		_, err := getAccessToken(username, password)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("test logging in using token", func(t *testing.T) {
		t.Parallel()
		token, _ := getAccessToken(username, password)
		returnedUsername, err := verifyToken(token)
		if err != nil {
			t.Error(err)
		} else if returnedUsername != username {
			t.Errorf("%s: returned username does not match the actual one", t.Name())
		}
	})
}
