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
package post

import (
	"github.com/spf13/cobra"
	"github.com/z-t-y/flogo/cmd"
)



var postLongDescription = `Publish, edit or view your posts

Basic Usage:
- flogo post publish
- flogo post edit
- flogo post view
Refer to their help pages for more details (flogo post publish --help)
`

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Publish, edit or view your posts",
	Long:  postLongDescription,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cmd.RootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
