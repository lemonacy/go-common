package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func init() {
	Viper = viper.New()
	Viper.SetConfigName("config")
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath(".")
	if err := Viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Viper.SetDefault("server.port", "8080")
}
