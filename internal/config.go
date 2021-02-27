package internal

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config is config model for the project.
type Config struct {
	// Scheduler check time (in minutes).
	Interval int `envconfig:"INTERVAL" default:"10"`
	// Local timezone.
	Timezone string `envconfig:"TIMEZONE" default:"UTC"`
	// Timezone location.
	Location *time.Location
	// Discord command prefix.
	Prefix string `envconfig:"PREFIX" default:"!"`
	// Discord token.
	Token string `envconfig:"TOKEN"`
	// MongoDB connection string.
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

	// Set timezone location.
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return cfg, ErrInvalidTimezone
	}

	cfg.Location = loc

	return cfg, nil
}
