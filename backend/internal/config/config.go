package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// Database
	DBHost     string `env:"POSTGRES_HOST"`
	DBPort     int    `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`

	// Cache
	CacheSize int `env:"CACHE_SIZE"`

	// HTTP
	HttpPort string `env:"HTTP_PORT"`
	// Environment
	AppEnv string `env:"APP_ENV"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

func init() {
	viper.SetConfigFile(".env")
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
}

func LoadConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read the .env file: %w", err)
	}

	config := &Config{
		DBHost:     viper.GetString("POSTGRES_ADDR"),
		DBPort:     viper.GetInt("POSTGRES_PORT"),
		DBUser:     viper.GetString("POSTGRES_USER"),
		DBPassword: viper.GetString("POSTGRES_PASSWORD"),
		DBName:     viper.GetString("POSTGRES_DB"),
		CacheSize:  viper.GetInt("CACHE_SIZE"),

		HttpPort: viper.GetString("HTTP_PORT"),
		AppEnv:   viper.GetString("APP_ENV"),
	}

	return config, nil
}
