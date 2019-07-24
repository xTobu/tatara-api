package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("toml")
	config.SetConfigName(env)
	config.AddConfigPath("../configs/")
	config.AddConfigPath("configs/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

// GetConfig : 取得 config by viper
func GetConfig() *viper.Viper {
	return config
}
