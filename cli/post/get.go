/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	md "github.com/JohannesKaufmann/html-to-markdown"
	mr "github.com/MichaelMure/go-term-markdown"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific post by id",
	Long:  `Get a specific post by id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("this command accepts only 1 arg, received", len(args))
			os.Exit(1)
		}
		postId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		post, err := GetPost(accessToken, postId)
		cobra.CheckErr(err)
		fmt.Println(u.Segmenter)
		fmt.Println("ID:         ", post.ID)
		fmt.Println("Author:     ", post.Author.Username)
		fmt.Println("Author ID:  ", post.Author.ID)
		fmt.Println("Title:      ", post.Title)
		fmt.Println("Private:    ", post.Private)
		fmt.Println("Columns:    ", post.Columns)
		fmt.Println("Comments:   ", post.Comments)
		fmt.Println("URL:        ", post.Self)
		fmt.Print("Content:      ")
		converter := md.NewConverter("", true, nil)
		markdown, err := converter.ConvertString(post.Content)
		cobra.CheckErr(err)
		content := mr.Render(markdown, 120, 0)
		fmt.Println(string(content))
	},
}

func GetPost(accessToken string, postId int) (post u.Post, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", u.URLFor("/api/v3/post/%d", postId), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = u.CheckStatusCode(resp, 200)
	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&post)
	return
}

func init() {
	postCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
