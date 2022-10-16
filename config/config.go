package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         int    `mapstructure:"DB_PORT"`
	DBUsername     string `mapstructure:"DB_USERNAME"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBDatabaseName string `mapstructure:"DB_DATABASE_NAME_RAW"`
}

// GetPGConnectionString: returns a postgres connection string withouth the database name
func (c *Config) GetPGConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword)
}

// GetDBConnectionString: returns a postgres connection string with the database name
func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBDatabaseName)
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	return config, err
}
