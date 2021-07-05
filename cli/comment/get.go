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
package comment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	. "github.com/z-t-y/flogo/utils"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get comment by id",
	Long:  `Get a comment by id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: you did not specify the id")
			os.Exit(1)
		}
		commentId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)
		accessToken, err := GetLocalAccessToken()
		cobra.CheckErr(err)
		comment, err := getComment(accessToken, commentId)
		cobra.CheckErr(err)
		fmt.Println("----------------------------------------")
		fmt.Println("ID:         ", comment.ID)
		fmt.Println("Author:     ", comment.Author.Username)
		fmt.Println("Author ID:  ", comment.Author.ID)
		fmt.Println("Body:       ", comment.Body)
		if comment.Replying != 0 {
			fmt.Println("Replying:   ", comment.Replying)
		}
	},
}

func getComment(accessToken string, commentId int) (comment Comment, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", URLFor("/api/v3/comment/%d", commentId), nil)
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
	json.NewDecoder(resp.Body).Decode(&comment)
	return
}

func init() {
	commentCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
