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
	FlogURL     string `mapstructure:"FLOG_URL"`
	AccessToken string `mapstructure:"ACCESS_TOKEN"`
}

func LoadConfig() (config Config, err error) {
	err = viper.Unmarshal(&config)
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
