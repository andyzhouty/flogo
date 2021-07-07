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
package column

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienroland/usg"
	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

var postIds []int
var name string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a column",
	Long:  `Create a column by specifying post ids`,
	Run: func(cmd *cobra.Command, args []string) {
		postIds := make([]int, len(args))
		for i, arg := range args {
			id, err := strconv.Atoi(arg)
			cobra.CheckErr(err)
			postIds[i] = id
		}
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		_, err = createColumn(accessToken, postIds, name)
		cobra.CheckErr(err)
	},
}

func createColumn(accessToken string, postIds []int, name string) (column u.Column, err error) {
	var data struct {
		Name  string `json:"name"`
		Posts []int  `json:"post_ids"`
	}
	data.Name = name
	data.Posts = postIds
	body, err := json.Marshal(data)
	cobra.CheckErr(err)

	client := http.Client{}
	req, err := http.NewRequest("POST", u.URLFor("/api/v3/column/create"), bytes.NewReader(body))
	cobra.CheckErr(err)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	cobra.CheckErr(err)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println(usg.Get.Tick, "Successfully added column", name)
	} else {
		err = u.CheckStatusCode(resp, 200)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&column)
	return
}

func init() {
	columnCmd.AddCommand(createCmd)
	createCmd.Flags().IntSliceVarP(&postIds, "posts", "p", []int{}, "Post ids")
	createCmd.Flags().StringVarP(&name, "name", "n", "", "Column name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
