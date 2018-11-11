package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func LoadConfig() {

	viper.BindEnv("app_env")
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	cURL := viper.GetString("consul_url")
	cPath := viper.GetString("consul_path")
	if cURL == "" {
		log.Fatal("CONSUL_URL missing")
	}
	if cPath == "" {
		log.Fatal("CONSUL_PATH missing")
	}

	viper.AddRemoteProvider("consul", cURL, cPath)
	viper.SetConfigType("yml")

	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatal(fmt.Sprintf(`%s named "%s"`, err.Error(), cPath))
	}
}
