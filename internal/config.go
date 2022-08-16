package internal

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config is config model for the project.
type Config struct {
	// Scheduler check time (in minutes).
	Interval int `envconfig:"INTERVAL" default:"10"`
	// Discord command prefix.
	Prefix string `envconfig:"PREFIX" default:"!"`
	// Discord token.
	Token string `envconfig:"TOKEN"`
	// MongoDB config.
	DB dbConfig `envconfig:"DB"`
}

type dbConfig struct {
	URI      string `envconfig:"URI"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
}

const envPath = "../../.env"
const envPrefix = "NXD"

// GetConfig to read and parse env.
func GetConfig() (cfg Config, err error) {
	// Load .env file if exist.
	godotenv.Load(envPath)

	// Convert env to struct.
	if err = envconfig.Process(envPrefix, &cfg); err != nil {
		return cfg, err
	}

	// Validate config.
	if cfg.Token == "" {
		return cfg, ErrRequiredToken
	}

	if len(cfg.Prefix) == 0 {
		return cfg, ErrRequiredPrefix
	}

	if cfg.Interval <= 0 {
		return cfg, ErrInvalidInterval
	}

	if cfg.DB.URI == "" {
		return cfg, ErrRequiredDB
	}

	return cfg, nil
}
