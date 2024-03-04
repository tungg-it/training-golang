package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ConfigAppEnv struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	Prefix        string `mapstructure:"PREFIX"`
	Host          string `mapstructure:"HOST"`
	Port          string `mapstructure:"PORT"`
	MySqlHost     string `mapstructure:"MYSQL_HOST"`
	MySqlPort     string `mapstructure:"MYSQL_PORT"`
	MySqlUser     string `mapstructure:"MYSQL_USER"`
	MySqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySqlDatabase string `mapstructure:"MYSQL_DATABASE"`
}

func LoadConfig() *ConfigAppEnv {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // Enable reading environment variables

	loadEnvIntoViper()

	// If the .env file exists, read it
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error reading config file: ", err)
		}
	}

	config := &ConfigAppEnv{}
	err := viper.Unmarshal(config)
	if err != nil {
		fmt.Println("Environment can't be loaded: ", err)
	}

	// Set default config if empty
	if config.Environment == "" {
		config.Environment = "development"
	}

	if config.Prefix == "" {
		config.Prefix = "api"
	}

	if config.Host == "" {
		config.Host = "0.0.0.0"
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.MySqlHost == "" {
		config.MySqlHost = "localhost"
	}

	if config.MySqlPort == "" {
		config.MySqlPort = "3306"
	}

	if config.MySqlUser == "" {
		config.MySqlUser = "root"
	}

	if config.MySqlPassword == "" {
		config.MySqlPassword = "root"
	}

	if config.MySqlDatabase == "" {
		config.MySqlDatabase = "db_dev"
	}

	return config
}

// monkey patching for loading current env into viper
func loadEnvIntoViper() {
	// Retrieve the environment variables
	envVars := os.Environ()

	// Iterate over each environment variable
	for _, env := range envVars {
		// Split the string into key and value
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]

		// Print the key and value
		viper.Set(key, value)
	}
}
