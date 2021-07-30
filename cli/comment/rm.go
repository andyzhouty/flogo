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
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a comment by id",
	Long:  `Remove a comment by id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("this command accepts 1 arg, received", len(args))
			os.Exit(1)
		}
		commentId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		err = RmComment(accessToken, commentId)
		cobra.CheckErr(err)
	},
}

func RmComment(accessToken string, commentId int) (err error) {
	client := http.Client{}
	req, err := http.NewRequest("DELETE", u.URLFor("/api/v3/comment/%d", commentId), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode == 403 {
		fmt.Println("Error: you don't have the permission to delete the comment")
		os.Exit(1)
	}
	err = u.CheckStatusCode(resp, 204)
	return
}

func init() {
	commentCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
