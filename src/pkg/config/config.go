package config

import (
	"PicusFinalCase/src/pkg/errorHandler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// LoadConfig It searches for yaml file in the same directory according to the
//incoming file name. It gives an error that the file cannot be found.
//Converts the file it finds to Config type.
func LoadConfig(filename string) (*Config, interface{}) {

	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("The file with the given filename could not be found at the specified address. : %v", err)
			return nil, errorHandler.ConfigNotFoundError
		}
		return nil, err
	}

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Error("Failed to convert read yaml file to config. : %v", err)
		return nil, errorHandler.UnmarshalError
	}

	return &c, nil
}
