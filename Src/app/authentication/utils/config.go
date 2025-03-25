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

// type ServerConfig struct {
// 	Localhost string `yaml:"localhost"`
//}

type Config struct {
	// Server   ServerConfig `yaml:"server"`
	Database DBConfig `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("applications")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")

	// viper.config := &Config{}
	// file, err := os.ReadFile("./resources/applications.yml")
	// if err != nil {
	// 	return nil, err
	// }
	// err = yaml.Unmarshal(file, config)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("Hello", config.Database.Username)
	// return config, nil

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	config := &Config{}

	// Unmarshal the config into the struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	// Print for debugging (optional)
	fmt.Println("Database Username:", config.Database.Username)

	return config, nil
}
