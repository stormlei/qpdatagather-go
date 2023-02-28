package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetEnvPrefix("qpdatagather")
	config.AutomaticEnv()
	config.SetConfigType("toml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file" + err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
