package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Load loads the configuration from the environment.
func Load() (Config, error) {
	config := Config{}

	err := godotenv.Load()
	if err != nil {
		// nothing major
	}

	err = envconfig.Process("PLUGIN", &config)

	// even if empty - no point in not setting it
	if config.DotscienceHost == "" {
		config.DotscienceHost = "https://cloud.dotscience.com"
	}

	return config, err
}

// MustLoad loads the configuration from the environment
// and panics if an error is encountered.
func MustLoad() Config {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	return config
}
