/*
Copyright © 2021 Andy Zhou

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package column

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/z-t-y/flogo/cli/auth"
	p "github.com/z-t-y/flogo/cli/post"
)

func TestCreateColumn(t *testing.T) {
	accessToken, err := auth.GetAccessToken(username, password)
	if err != nil {
		t.Error(err)
	}
	posts := make([]int, 5)
	for i := 0; i < 5; i++ {
		title := "Flogo Post Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
		content := title
		post, err := p.UploadPost(title, content, accessToken)
		if err != nil {
			t.Error(err)
		}
		posts[i] = post.ID
	}
	name := "Flogo Column Unit Test - " + strconv.FormatInt(time.Now().Unix(), 10)
	c, err := CreateColumn(accessToken, posts, name)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c.Name)
	if c.Name != name {
		t.Errorf("TestCreateColumn: expect column name %s, actual %s", name, c.Name)
	}
}
