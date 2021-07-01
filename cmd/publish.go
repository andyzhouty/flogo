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
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienroland/usg"
	. "github.com/z-t-y/flogo/utils"

	"github.com/gomarkdown/markdown"
	"github.com/spf13/cobra"
)

var private bool
var markdownFile, htmlFile, postTitle string

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish your post in MarkDown or HTML",
	Long: `Publish your post in MarkDown or HTML.

Usage:
- flogo post publish (<post-name>) -m/--markdown <markdown-file> (-c <column-name>) (-p/--private)
- flogo post publish (<post-name>) --html <html-file> (-c <column-name>) (-p/--private)

Note:
1. <!doctype html>, <head>, <body> tags are not necessary in the html, just post your content!
2. You are recommended to use markdown since it's simpler and will cause less problems while uploading.
3. If you use
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if markdownFile != "" && htmlFile != "" {
			fmt.Println("You should not use BOTH of the --markdown and the --html flags.")
			os.Exit(1)
		}
		var htmlContent string
		if markdownFile != "" {
			content, err := ioutil.ReadFile(markdownFile)
			cobra.CheckErr(err)
			htmlContent = string(markdown.ToHTML(content, nil, nil))
		} else if htmlFile != "" {
			content, err := ioutil.ReadFile(htmlFile)
			cobra.CheckErr(err)
			htmlContent = string(content)
		} else {
			fmt.Println("Error: input file does not exist or invalid")
			os.Exit(1)
		}
		if postTitle == "" {
			fmt.Println("Error: post title empty")
			os.Exit(1)
		}
		accessToken, err := GetAccessToken()
		cobra.CheckErr(err)
		_, err = uploadPost(postTitle, htmlContent, accessToken)
		cobra.CheckErr(err)
	},
}

func uploadPost(postTitle string, htmlContent string, accessToken string) (post Post, err error) {
	flogURL, err := GetFlogURL()
	cobra.CheckErr(err)
	var data struct {
		Title       string `json:"title"`
		HTMLContent string `json:"content"`
	}
	data.Title = postTitle
	data.HTMLContent = htmlContent
	body, err := json.Marshal(data)
	cobra.CheckErr(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", flogURL+"/api/v3/post/add", bytes.NewReader(body))
	cobra.CheckErr(err)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := client.Do(req)
	cobra.CheckErr(err)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println(usg.Get.Tick, "Successfully added post", postTitle)
	} else {
		err = CheckStatusCode(resp, 200)
		return
	}
	respBody := make([]byte, resp.ContentLength)
	resp.Body.Read(respBody)
	err = json.Unmarshal(respBody, &post)
	return
}

func init() {
	postCmd.AddCommand(publishCmd)
	publishCmd.Flags().BoolVarP(&private, "private", "p", false, "Set post privacy")
	publishCmd.Flags().StringVarP(&markdownFile, "markdown", "m", "", "Post written in MarkDown.")
	publishCmd.Flags().StringVar(&htmlFile, "html", "", "Post written in HTML.") // avoid conflict with --help
	publishCmd.Flags().StringVarP(&postTitle, "title", "t", "", "Post title.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
