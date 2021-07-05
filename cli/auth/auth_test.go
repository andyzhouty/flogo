package auth

import (
	"github.com/z-t-y/flogo/utils"
	"testing"
)

func TestAuth(t *testing.T) {
	t.Parallel()
	utils.SetArgs("auth")
	utils.SetArgs("auth", "login")
}
