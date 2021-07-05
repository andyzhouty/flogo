package comment

import (
	"github.com/z-t-y/flogo/utils"
	"os"
	"testing"
)

var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func TestCommentCommand(t *testing.T) {
	t.Parallel()
	err := utils.SetArgs("comment").Run()
	if err != nil {
		t.Error(err)
	}
}