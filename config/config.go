package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresDBUsername string `mapstructure:"POSTGRES_DB_USERNAME"`
	PostgresDBPassword string `mapstructure:"POSTGRES_DB_PASSWORD"`
	PostgresDBHost     string `mapstructure:"POSTGRES_DB_HOST"`
	PostgresDBName     string `mapstructure:"POSTGRES_DB_NAME"`
	MigratePath        string `mapstructure:"MIGRATE_PATH"`

	ServerHost string `mapstructure:"SERVER_HOST"`

	ApiKey        string `mapstructure:"API_KEY"`
	ApiURI        string `mapstructure:"API_URI"`
	ClientTimeout int    `mapstructure:"CLIENT_TIMEOUT"`
}

func New() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("read in config failed: %w", err)
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshal failed: %w", err)
	}
	return config, nil
}

func (c Config) GetDBUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", c.PostgresDBUsername, c.PostgresDBPassword, c.PostgresDBHost, c.PostgresDBName)
}
