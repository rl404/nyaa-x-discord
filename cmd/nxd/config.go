package main

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
	"github.com/rl404/nyaa-x-discord/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type config struct {
	Discord  discordConfig  `envconfig:"DISCORD"`
	DB       dbConfig       `envconfig:"DB"`
	Cron     cronConfig     `envconfig:"CRON"`
	Log      logConfig      `envconfig:"LOG"`
	Newrelic newrelicConfig `envconfig:"NEWRELIC"`
}

type discordConfig struct {
	Token  string `envconfig:"TOKEN" validate:"required"`
	Prefix string `envconfig:"PREFIX" validate:"required" mod:"default=!"`
}

type dbConfig struct {
	URI      string `envconfig:"URI" validate:"required"`
	Name     string `envconfig:"NAME" validate:"required" mod:"default=nyaaXdiscord"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
}

type logConfig struct {
	Level utils.LogLevel `envconfig:"LEVEL" default:"-1"`
	JSON  bool           `envconfig:"JSON" default:"false"`
	Color bool           `envconfig:"COLOR" default:"true"`
}

type newrelicConfig struct {
	Name       string `envconfig:"NAME" default:"nxd"`
	LicenseKey string `envconfig:"LICENSE_KEY"`
}

type cronConfig struct {
	Interval time.Duration `envconfig:"INTERVAL" validate:"required,gt=0" mod:"default=10m"`
}

const envPath = "../../.env"
const envPrefix = "NXD"

func getConfig() (*config, error) {
	var cfg config

	// Load .env file.
	_ = godotenv.Load(envPath)

	// Convert env to struct.
	if err := envconfig.Process(envPrefix, &cfg); err != nil {
		return nil, err
	}

	// Validate.
	if err := utils.Validate(&cfg); err != nil {
		return nil, err
	}

	// Init global log.
	utils.InitLog(cfg.Log.Level, cfg.Log.JSON, cfg.Log.Color)

	return &cfg, nil
}

func newDB(cfg dbConfig) (*mongo.Database, error) {
	nrMongo := nrmongo.NewCommandMonitor(nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start connection.
	client, err := mongo.Connect(ctx, options.
		Client().
		ApplyURI(cfg.URI).
		SetAuth(options.Credential{
			Username: cfg.User,
			Password: cfg.Password,
		}).
		SetMonitor(nrMongo))
	if err != nil {
		return nil, err
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	// Ping test.
	if err = client.Ping(ctx2, nil); err != nil {
		return nil, err
	}

	return client.Database(cfg.Name), nil
}
