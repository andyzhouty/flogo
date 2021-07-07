package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/oleiade/reflections"

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

func WriteDefault() error {
	fields, err := reflections.Fields(DefaultConfig)
	if err != nil {
		return err
	}
	for _, field := range fields {
		value, err := reflections.GetField(DefaultConfig, field)
		if err != nil {
			return err
		}
		tags, err := reflections.Tags(DefaultConfig, "mapstructure")
		if err != nil {
			return err
		}
		for fieldName, tagValue := range tags {
			if fieldName == field {
				if value.(string) != "" {
					err = WriteToConfig(tagValue, value.(string))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return err
}

func GetConfig(configName string) (value interface{}, err error) {
	var config Config
	viper.Unmarshal(&config)
	tags, err := reflections.Tags(config, "mapstructure")
	for fieldName, tagValue := range tags {
		if tagValue == configName {
			value, err = reflections.GetField(config, fieldName)
			if value == nil {
				value, err = reflections.GetField(DefaultConfig, fieldName)
			}
		}
	}
	return
}

func GetLocalAccessToken() (accessToken string, err error) {
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
