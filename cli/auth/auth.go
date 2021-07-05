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
	"fmt"
	"github.com/z-t-y/flogo/cmd"

	"github.com/spf13/cobra"
)

var longHelp = `A command that authenticates the current user, using bearer tokens.
Note that the bearer tokens expire EVERY YEAR, so during the year you should not share your tokens to others.
`

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "A command that authenticates the current user",
	Long:  longHelp,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(longHelp)
	},
}

func init() {
	cmd.RootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
