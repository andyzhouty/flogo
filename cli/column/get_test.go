package column

import (
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
	p "github.com/z-t-y/flogo/cli/post"
)

func TestGetColumn(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	title := "Flogo Post Unit Test - " + strconv.FormatInt(time.Now().UnixNano(), 10)
	content := title
	postIds := make([]int, 5)
	for i := 0; i < 5; i++ {
		post, err := p.UploadPost(title, content, accessToken)
		if err != nil {
			return
		}
		postIds[i] = post.ID
	}
	name := "Flogo Column Unit Test - " + strconv.FormatInt(time.Now().UnixNano(), 10)
	column, err := CreateColumn(accessToken, postIds, name)
	if err != nil {
		t.Error(err)
	}
	if column.Name != name {
		t.Errorf("TestGetColumn: expected column name %s, actual %s", name, column.Name)
	}
}
