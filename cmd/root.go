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
	"errors"
	"io/ioutil"
	"os"
	"reflect"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/z-t-y/flogo/utils"
)

var cfgFile string

var ErrInvalidToken = errors.New("token invalid")

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "flogo",
	Short: "The Flog CLI",
	Long:  `A simple command line interface for accessing Flog from your terminal!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flogo)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		cfgFile = home + "/.flogo"
		// Search config in home directory with name ".flogo" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".flogo")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	_, err := os.Stat(cfgFile)
	if err != nil {
		os.Create(cfgFile)
		utils.WriteDefault()
	}
	content, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		cobra.CheckErr(err)
	}
	if reflect.DeepEqual(content, []byte{10}) { // if file content is empty
		utils.WriteDefault()
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(err)
	}
}
