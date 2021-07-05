package post

import (
	"github.com/z-t-y/flogo/utils"
	"os"
	"testing"
)

var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func TestPost(t *testing.T) {
	t.Parallel()
	utils.SetArgs("post")
}