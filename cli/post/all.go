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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all public posts and your private posts",
	Long: `List all public posts and your private posts
By default, it'll show you the id, title of the post and whether it is private.

Note:
You should not use the head and tail options together. It may lead to unexpected results.
`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		posts, err := GetPostsFrom("/api/v3/post/all", accessToken)
		cobra.CheckErr(err)
		if tail > 0 && head > 0 {
			fmt.Println("Error: you cannot use the head and tail options together")
			os.Exit(1)
		}
		if tail > 0 {
			posts = posts[len(posts)-int(tail):]
		} else if head > 0 {
			posts = posts[:head]
		}
		for _, post := range posts {
			u.OutputPost(post, short, verbose, veryVerbose)
		}
	},
}

func init() {
	postCmd.AddCommand(allCmd)
	allCmd.Flags().BoolVarP(&short, "short", "s", false, "Short Mode(only showing the IDs)")
	allCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Mode(showing the column(s) and first 200 characters of each post)")
	allCmd.Flags().BoolVarP(&veryVerbose, "very-verbose", "V", false,
		"Very Verbose Mode(showing the whole content, URL, column(s) and comment(s) of each post)")
	allCmd.Flags().Uint32VarP(&head, "head", "H", 0, "Show the first H posts")
	allCmd.Flags().Uint32VarP(&tail, "tail", "t", 0, "Return the newest posts, defaulting to 0")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
