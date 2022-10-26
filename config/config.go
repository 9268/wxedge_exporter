package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port      int      `json:"port"`
	Host      []string `json:"host"`
	NameSpace []string `json:"namespace"`
}

var Configs Config

func ReadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")

	viper.SetDefault("host", []string{"http://127.0.0.1:18888"})
	viper.SetDefault("port", 9001)
	viper.SetDefault("namespace", "wxedge")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Configs.Port = viper.GetInt("port")
	Configs.Host = viper.GetStringSlice("host")
	Configs.NameSpace = viper.GetStringSlice("namespace")
}
func GetHost() []string {
	return viper.GetStringSlice("Host")
}
