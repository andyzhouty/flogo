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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/z-t-y/flogo/utils"
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
		if accessToken == "" {
			useUsernamePassword()
		}
	},
}

func useUsernamePassword() {
	var username, password string
	fmt.Print("Enter your flog username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)
	token, err := getAccessToken(username, password)
	if err != nil {
		cobra.CheckErr(err)
	}
	viper.Set("access_token", token)
	fmt.Println(viper.AllSettings())
	err = utils.WriteToConfig()
	cobra.CheckErr(err)
}

func getAccessToken(username string, password string) (string, error) {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)
	config, err := utils.LoadConfig()
	if err != nil {
		return "", err
	}
	flogURL := config.FlogURL
	if flogURL == "" {
		flogURL = utils.DefaultConfig.FlogURL
	}

	resp, err := http.PostForm(flogURL+"/v3/token", data)
	if err != nil {
		return "", err
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 400 {
		return "", errors.New("invalid username or password")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var t *utils.TokenResp
	json.Unmarshal(body, &t)
	return t.AccessToken, err
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
