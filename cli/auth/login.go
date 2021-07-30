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
package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/z-t-y/flogo/cmd"

	"github.com/julienroland/usg"
	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

var accessToken string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		var err error
		if accessToken == "" {
			var password string
			fmt.Print("Enter your flog Username: ")
			fmt.Scanln(&username)
			fmt.Print("Enter your Password: \033[8m") // hide the input
			fmt.Scanln(&password)
			fmt.Print("\033[28m") // show the input
			err = useUsernamePassword(username, password)
			cobra.CheckErr(err)
		} else {
			username, err = verifyToken(accessToken)
			cobra.CheckErr(err)
		}
		fmt.Println(usg.Get.Tick, "Logined as", username)
	},
}

func useUsernamePassword(username, password string) (err error) {
	token, err := GetAccessToken(username, password)
	if err != nil {
		return
	}
	err = u.WriteToConfig("access_token", token)
	return
}

func GetAccessToken(username string, password string) (string, error) {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	resp, err := http.PostForm(u.URLFor("/api/v3/token"), data)
	cobra.CheckErr(err)
	if resp.StatusCode == 400 {
		return "", errors.New("invalid username or password")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var t u.Token
	json.Unmarshal(body, &t)
	return t.AccessToken, err
}

func verifyToken(token string) (username string, err error) {
	data := url.Values{}
	data.Add("token", token)
	resp, err := http.PostForm(u.URLFor("/api/v3/token/verify"), data)
	if err != nil {
		return
	}
	if resp.StatusCode == 401 {
		return "", cmd.ErrInvalidToken
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var schema struct {
		Username string `json:"Username"`
		Valid    bool   `json:"valid"`
	}
	json.Unmarshal(body, &schema)
	if !schema.Valid {
		err = cmd.ErrInvalidToken
	}
	username = schema.Username
	return
}

func init() {
	authCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&accessToken, "token", "t", "", "Access Token")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
