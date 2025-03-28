package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	Database DBConfig `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("applications")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}


	return config, nil
}
