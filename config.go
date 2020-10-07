package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// cfg is config data.
var cfg Config

// Config is config model for the project.
type Config struct {
	// Scheduler check time (in minutes).
	Interval int
	// Local timezone.
	Timezone string
	// Timezone location.
	Location *time.Location
	// Discord token.
	Token string
	// MongoDB connection string.
	DB string
}

const (
	// EnvPath is .env file path.
	EnvPath = ".env"
	// EnvPrefix is environment
	EnvPrefix = "NXD"
	// DefaulInterval is default scheduler check time (5 minutes).
	DefaulInterval = 5
	// DefaultTimezone is default timezone.
	DefaultTimezone = "UTC"
)

// setConfig to read .env and set the config.
func setConfig() error {
	// Set default config.
	cfg.Interval = DefaulInterval
	cfg.Timezone = DefaultTimezone

	// Load .env file if exist.
	godotenv.Load(EnvPath)

	// Convert env to struct.
	envconfig.Process(EnvPrefix, &cfg)

	// Validate config.
	if cfg.Token == "" {
		return ErrRequiredToken
	}

	if cfg.Interval <= 0 {
		return ErrInvalidInterval
	}

	if cfg.DB == "" {
		return ErrRequiredDB
	}

	// Set timezone location.
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return ErrInvalidTimezone
	}

	cfg.Location = loc

	// Set db connection.
	err = setConnection()
	if err != nil {
		return err
	}

	return nil
}
