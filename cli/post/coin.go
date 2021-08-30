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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

var postId, amount int

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "Give your coins to a post",
	Long:  `Give your coins to a post.`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		if amount < 1 || amount > 2 {
			fmt.Println("Error: invalid amount of coins")
			os.Exit(1)
		}
		post, err := GiveCoin(postId, amount, accessToken)
		cobra.CheckErr(err)
		u.OutputPost(post, false, false, true)
	},
}

func GiveCoin(postId, amount int, accessToken string) (post u.Post, err error) {
	var data struct {
		Amount int `json:"amount"`
	}
	data.Amount = amount
	body, err := json.Marshal(data)
	cobra.CheckErr(err)

	client := &http.Client{}
	req, err := http.NewRequest("POST", u.URLFor("/api/v3/post/coin/%d", postId), bytes.NewReader(body))
	cobra.CheckErr(err)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	cobra.CheckErr(err)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = u.CheckStatusCode(resp, 200)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&post)
	return
}

func init() {
	postCmd.AddCommand(coinCmd)

	coinCmd.Flags().IntVarP(&amount, "amount", "a", 1, "Amount of coins to give")
	coinCmd.Flags().IntVarP(&postId, "post", "p", 0, "Post id to give coins to")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
