package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (Config, error) {
	var c Config

	viper.AddConfigPath("./pkg/config")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return c, err
	}

	err = viper.Unmarshal(&c)

	return c, err
}
