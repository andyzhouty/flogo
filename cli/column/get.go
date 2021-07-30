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
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a column",
	Long:  `Get a column specified by id`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("this command accepts 1 arg, received", len(args))
		}
		columnId, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		column, err := GetColumn(accessToken, columnId)
		cobra.CheckErr(err)
		// Output the column into the terminal.
		fmt.Println(u.Segmenter)
		fmt.Println("ID:         ", column.ID)
		fmt.Println("Author:     ", column.Author.Username)
		fmt.Println("Author ID:  ", column.Author.ID)
		fmt.Println("Name:       ", column.Name)
		fmt.Print("Post IDs:   ")
		for _, post := range column.Posts {
			fmt.Print(post.ID, ", ")
		}
		fmt.Println()
		fmt.Println("URL:        ", column.URL)
	},
}

func GetColumn(accessToken string, columnId int) (column u.Column, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", u.URLFor("/api/v3/column/%d", columnId), nil)
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
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&column)
	return
}

func init() {
	columnCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
