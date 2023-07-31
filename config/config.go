package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Driver   string
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	}
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("wms.yaml")   // name of wms.yaml file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the wms.yaml file does not have the extension in the name
	viper.AddConfigPath("$HOME/.wms") // call multiple times to add many search paths
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read wms.yaml file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal wms.yaml: %v", err)
	}
	return &config, nil
}
