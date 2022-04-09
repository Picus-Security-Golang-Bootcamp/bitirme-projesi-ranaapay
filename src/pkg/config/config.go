package config

import (
	"PicusFinalCase/src/pkg/errorHandler"
	"github.com/spf13/viper"
)

func LoadConfig(filename string) (*Config, interface{}) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	viper.SetConfigType("yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errorHandler.ConfigNotFoundError
		}
		return nil, err
	}
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		return nil, errorHandler.UnmarshalError
	}
	return &c, nil
}
