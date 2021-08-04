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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienroland/usg"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

var postId, replyingId int
var content string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add your comment",
	Long:  `Add your comment to a post specified by post id and replying comment id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if postId == 0 {
			fmt.Println(usg.Get.Cross, "You should specify a post you want to comment.")
			os.Exit(1)
		}
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		_, err = AddComment(accessToken, content, postId, replyingId)
		cobra.CheckErr(err)
	},
}

func AddComment(accessToken string, content string, postId int, replyingId int) (comment u.Comment, err error) {
	data := map[string]interface{}{}
	data["body"] = content
	data["post_id"] = postId
	if replyingId != 0 {
		data["reply_id"] = replyingId
	}
	fmt.Println(data)
	body, err := json.Marshal(data)
	cobra.CheckErr(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", u.URLFor("/api/v3/comment/add"), bytes.NewReader(body))
	cobra.CheckErr(err)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	cobra.CheckErr(err)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("%s Successfully added comment <Comment %s>\n", usg.Get.Tick, content)
	} else {
		err = u.CheckStatusCode(resp, 200)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&comment)
	return
}

func init() {
	commentCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&postId, "post", "p", 0, "The id of the post")
	addCmd.Flags().IntVarP(&replyingId, "reply", "r", 0, "The id of the replying comment")
	addCmd.Flags().StringVarP(&content, "content", "c", "", "The content of your comment")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
