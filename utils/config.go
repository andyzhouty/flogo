package utils

import (
	"encoding/json"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DefaultConfig = Config{
	FlogURL: "http://flog-web.herokuapp.com/api",
}

type Config struct {
	FlogURL     string `mapstructure:"flog_url"`
	AccessToken string `mapstructure:"access_token"`
}

func LoadConfig() (config Config, err error) {
	err = viper.Unmarshal(&config)
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

func WriteToConfig() (err error) {
	config := viper.AllSettings()
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
