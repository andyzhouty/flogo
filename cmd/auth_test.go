package cmd

import "testing"

func TestAuth(t *testing.T) {
	t.Parallel()
	testArgs("auth")
	testArgs("auth", "login")
}
