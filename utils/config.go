package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oleiade/reflections"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DefaultConfig = Config{
	FlogURL: "https://flog-web.herokuapp.com",
}

type Config struct {
	FlogURL     string `mapstructure:"flog_url"`
	AccessToken string `mapstructure:"access_token"`
}

func GetConfig(configName string) (value interface{}, err error) {
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	tags, err := reflections.Tags(config, "mapstructure")
	if err != nil {
		return
	}
	for fieldName, tagValue := range tags {
		if tagValue == configName {
			value, err = reflections.GetField(config, fieldName)
			if value == "" {
				value, err = reflections.GetField(DefaultConfig, fieldName)
			}
		}
	}
	return
}

func GetAccessToken() (accessToken string, err error) {
	var config Config
	err = viper.Unmarshal(&config)
	accessToken = config.AccessToken
	if accessToken == "" {
		err = errors.New("not logged in")
		return
	}
	return
}

func GetFlogURL() (flogURL string, err error) {
	var config Config
	err = viper.Unmarshal(&config)
	flogURL = config.FlogURL
	if flogURL == "" {
		flogURL = DefaultConfig.FlogURL
	}
	return
}

func WriteToConfig(key string, value string) (err error) {
	config := viper.AllSettings()
	config[key] = value
	fmt.Println(key, value)
	b, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return
	}
	home, err := homedir.Dir()
	cobra.CheckErr(err)
	f, err := os.Create(home + "/.flogo")
	if err != nil {
		return
	}
	f.WriteString(string(b))
	return
}
