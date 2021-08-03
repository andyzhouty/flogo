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
package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

var verbose, veryVerbose, short bool
var head, tail uint32

// selfCmd represents the self command
var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "List all your posts",
	Long: `List all your posts
By default, it'll show you the id, title of the post and whether it is private.

Note:
You should not use the head and tail options together. It may lead to unexpected results.
`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		posts, err := GetPostsFrom("/api/v3/self/posts", accessToken)
		cobra.CheckErr(err)
		if tail > 0 && head > 0 {
			fmt.Println("You cannot use the head and tail options together.")
			os.Exit(1)
		}
		if tail > 0 {
			posts = posts[len(posts)-int(tail):]
		} else if head > 0 {
			posts = posts[:head]
		}
		u.OutputPosts(posts, short, verbose, veryVerbose)
	},
}

func GetPostsFrom(url, accessToken string) (posts []u.Post, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.URLFor(url), nil)
	if err != nil {
		return
	}
	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = u.CheckStatusCode(resp, 200)

	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&posts)
	return
}

func init() {
	postCmd.AddCommand(selfCmd)
	selfCmd.Flags().BoolVarP(&short, "short", "s", false, "Short Mode(only showing the IDs)")
	selfCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Mode(showing the column(s) and first 200 characters of each post)")
	selfCmd.Flags().BoolVarP(&veryVerbose, "very-verbose", "V", false,
		"Very Verbose Mode(showing the whole content, URL, column(s) and comment(s) of each post)")
	selfCmd.Flags().Uint32VarP(&head, "head", "H", 0, "Show the first H posts")
	selfCmd.Flags().Uint32VarP(&tail, "tail", "t", 0, "Return the newest posts, defaulting to 0")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
