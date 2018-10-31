package config

import (
	"fmt"
	"github.com/spf13/viper"
)


var AppConfig *viper.Viper

func LoadConfig()  {
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.AddConfigPath(".")
	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	AppConfig = v
}
