package services

import (
	"github.com/spf13/viper"
	"strings"
)

var Config = viper.New()

func init() {

	Config.AutomaticEnv()
	Config.SetEnvPrefix("ENV")
	Config.SetEnvKeyReplacer(
		strings.NewReplacer(".", "_"),
	)
	Config.SetConfigFile("configs/config.yml")
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil {
		Logger.Error("cannot read config")
	}

}
