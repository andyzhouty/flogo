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
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	. "github.com/z-t-y/flogo/utils"
	"html"
	"net/http"
	"os"
	"strconv"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("this command accepts only 1 arg, received", len(args))
			os.Exit(1)
		}
		postId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)
		accessToken, err := GetAccessToken()
		cobra.CheckErr(err)
		post, err := getPost(accessToken, postId)
		cobra.CheckErr(err)
		fmt.Println("----------------------------------------")
		fmt.Println("Post ID:    ", post.ID)
		fmt.Println("Post title: ", post.Title)
		fmt.Println("Private:    ", post.Private)
		fmt.Println("Columns:    ", post.Columns)
		fmt.Println("Comments:   ", post.Comments)
		fmt.Println("URL:        ", post.Self)
		fmt.Println("Content:    ", html.UnescapeString(post.Content))
	},
}

func getPost(accessToken string, postId int) (post Post, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", URLFor("/api/v3/post/%d", postId), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = CheckStatusCode(resp, 200)
	if err != nil {
		return
	}
	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)
	err = json.Unmarshal(data, &post)
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
