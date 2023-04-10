package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.BindEnv("PORT")
	viper.BindEnv("DB_URL")

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	if config.DBUrl == "" && config.Port == "" {
		err = viper.ReadInConfig()
	}

	return
}
