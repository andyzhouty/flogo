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
	"fmt"

	"github.com/spf13/cobra"
)

var longDescription = `Display or change configuration settings for flogo.

Available configs:
- flog_url       Choose WHICH flog you use (default: https://flog-web.herokuapp.com)
                 You can also set it to the flog you deployed like http://localhost:5000.
- access_token   The flog access token.

Usage:
- flogo config get <key>  // Get a config from your config file (default $HOME/.flogo)
- flogo config set <key>  // Write a config to your config file
`

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display or change configuration settings for flogo.",
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(longDescription)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
