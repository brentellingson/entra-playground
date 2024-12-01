package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// OAuthConfiguration struct to hold the OAuth configuration values
type OAuthConfiguration struct {
	ClientID              string
	ClientSecret          string
	AuthorizationEndpoint string
	TokenEndpoint         string
	Scope                 string
	JwksURI               string
	Issuer                string
	Audience              string
}

// Config struct to hold all the configuration values
type Config struct {
	WebAPIA OAuthConfiguration
}

// NewConfig function to load the configuration values from the config file
func NewConfig() *Config {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
	viper.SetConfigFile("./config/config.secrets.yaml")
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}
	return &cfg
}
