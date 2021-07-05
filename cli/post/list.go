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
	md "github.com/JohannesKaufmann/html-to-markdown"
	mr "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"
	. "github.com/z-t-y/flogo/utils"
	"net/http"
	"strings"
)

var verbose, veryVerbose, short bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your posts",
	Long: `List all your posts
By default, it'll show you the id, title of the post and whether it is private.
`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := GetLocalAccessToken()
		cobra.CheckErr(err)
		posts, err := GetPosts(accessToken)
		cobra.CheckErr(err)
		converter := md.NewConverter("", true, nil)

		for _, post := range posts {
			switch {
			case short:
				fmt.Print(post.ID, " ")
			case verbose:
				fmt.Println("----------------------------------------")
				fmt.Println("Post ID:    ", post.ID)
				fmt.Println("Post title: ", post.Title)
				fmt.Println("Private:    ", post.Private)
				fmt.Println("Column(s):    ", post.Columns)
				fmt.Print("Content: ")
				if !(strings.Contains(post.Content, "</iframe>") || strings.Contains(post.Content, "<img")) {
					markdown, err := converter.ConvertString(post.Content)
					cobra.CheckErr(err)
					content := mr.Render(markdown, 120, 0)
					strContent := string(content)
					if len(strContent) > 200 {
						fmt.Println(strContent[:200])
					} else {
						fmt.Println(strContent)
					}
				} else {
					fmt.Println("Post Content contains iframes or images which cannot printed in the terminal")
					flogURL, err := GetFlogURL()
					cobra.CheckErr(err)
					fmt.Printf("Please visit %s%s to view it\n", flogURL, post.Self)
				}
			case veryVerbose:
				fmt.Println("----------------------------------------")
				fmt.Println("Post ID:    ", post.ID)
				fmt.Println("Post title: ", post.Title)
				fmt.Println("Private:    ", post.Private)
				fmt.Println("Columns:    ", post.Columns)
				fmt.Println("Comments:   ", post.Comments)
				fmt.Println("URL:        ", post.Self)
				fmt.Print("Content: ")
				if !(strings.Contains(post.Content, "</iframe>") || strings.Contains(post.Content, "<img")) {
					markdown, err := converter.ConvertString(post.Content)
					cobra.CheckErr(err)
					content := mr.Render(markdown, 120, 0)
					fmt.Println(string(content))
				} else {
					fmt.Println("Post Content contains iframes or images which cannot printed in the terminal")
					flogURL, err := GetFlogURL()
					cobra.CheckErr(err)
					fmt.Printf("Please visit %s%s to view it\n", flogURL, post.Self)
				}
			default:
				fmt.Println("----------------------------------------")
				fmt.Println("Post ID:    ", post.ID)
				fmt.Println("Post title: ", post.Title)
				fmt.Println("Private:    ", post.Private)
			}
		}
	},
}

func GetPosts(accessToken string) (posts []Post, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URLFor("/api/v3/self/posts"), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = CheckStatusCode(resp, 200)

	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&posts)
	return
}

func init() {
	postCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&short, "short", "s", false, "Short Mode(only showing the IDs)")
	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Mode(showing the column(s) and first 200 characters of each post)")
	listCmd.Flags().BoolVarP(&veryVerbose, "very-verbose", "V", false,
		"Very Verbose Mode(showing the whole content, URL, column(s) and comment(s) of each post)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
