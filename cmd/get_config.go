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
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/z-t-y/flogo/utils"
	"os"
)

// getConfigCmd represents the getConfig command
var getConfigCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your config",
	Long:  `Get the specified config, for all available configs, refer to 'flogo config --help'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("accepts 1 arg, received", len(args))
			os.Exit(1)
		}
		if value := viper.AllSettings()[args[0]]; value == nil {
			value, _ := utils.GetConfig(args[0])
			//cobra.CheckErr(err)
			if value == "" {
				fmt.Println("config key does not exist or not properly set")
				os.Exit(1)
			}
			fmt.Println(value)
		} else {
			fmt.Println(value)
		}
	},
}

func init() {
	configCmd.AddCommand(getConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
