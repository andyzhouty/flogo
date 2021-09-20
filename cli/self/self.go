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
package self

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/z-t-y/flogo/cmd"
	u "github.com/z-t-y/flogo/utils"
)

// selfCmd represents the self command
var selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Show your profile",
	Long:  `This command shows your profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		user, err := GetSelf(accessToken)
		cobra.CheckErr(err)
		fmt.Println("Username     :", user.Username)
		fmt.Println("Real Name    :", user.Name)
		fmt.Println("ID           :", user.ID)
		fmt.Println("Location     :", user.Location)
		fmt.Println("About Me     :", user.AboutMe)
		fmt.Println("Coins        :", user.Coins)
		fmt.Println("Experience   :", user.Experience)
		fmt.Println("Last Seen    :", user.LastSeen.Time)
		fmt.Println("Member Since :", user.MemberSince.Time)
	},
}

func GetSelf(accessToken string) (user u.User, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", u.URLFor("/api/v3/self"), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = u.CheckStatusCode(resp, 200)
	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&user)
	return
}

func init() {
	cmd.RootCmd.AddCommand(selfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
