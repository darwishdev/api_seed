package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	State             string `mapstructure:"STATE"`
	DBUser            string `mapstructure:"DB_USER"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBName            string `mapstructure:"DB_NAME"`
	TestDBName        string `mapstructure:"TEST_DB_NAME"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	TestDBSource      string `mapstructure:"TEST_DB_SOURCE"`
	FiberPort         string `mapstructure:"FIBER_PORT"`
	FiberHost         string `mapstructure:"FIBER_HOST"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	SupaUrl           string `mapstructure:"SUPA_URL"`
	SupaKey           string `mapstructure:"SUPA_KEY"`
}

func loadAncScan(config *Config) (err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("state")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = loadAncScan(&config)
	if err != nil {
		return
	}

	stateEnvFilePath := fmt.Sprintf("%s.env", config.State)
	viper.SetConfigName(stateEnvFilePath)
	err = loadAncScan(&config)
	if err != nil {
		return
	}

	viper.SetConfigName("shared.env")
	err = loadAncScan(&config)
	if err != nil {
		return
	}
	return
}
