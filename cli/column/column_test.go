package column

import (
	"os"
	"testing"

	"github.com/z-t-y/flogo/utils"
)

var username = os.Getenv("FLOG_USERNAME")
var password = os.Getenv("FLOG_PASSWORD")

func TestColumnCmd(t *testing.T) {
	err := utils.SetArgs("column").Run()
	if err != nil {
		t.Error(err)
	}
}
