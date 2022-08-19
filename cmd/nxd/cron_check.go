package main

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	_nr "github.com/rl404/fairy/log/newrelic"
	"github.com/rl404/nyaa-x-discord/internal/delivery/cron"
	discordRepository "github.com/rl404/nyaa-x-discord/internal/domain/discord/repository"
	discordClient "github.com/rl404/nyaa-x-discord/internal/domain/discord/repository/client"
	nyaaRepository "github.com/rl404/nyaa-x-discord/internal/domain/nyaa/repository"
	nyaaRSS "github.com/rl404/nyaa-x-discord/internal/domain/nyaa/repository/rss"
	templateRepository "github.com/rl404/nyaa-x-discord/internal/domain/template/repository"
	templateClient "github.com/rl404/nyaa-x-discord/internal/domain/template/repository/client"
	userRepository "github.com/rl404/nyaa-x-discord/internal/domain/user/repository"
	userDB "github.com/rl404/nyaa-x-discord/internal/domain/user/repository/db"
	"github.com/rl404/nyaa-x-discord/internal/service"
	"github.com/rl404/nyaa-x-discord/internal/utils"
)

func cronCheck() error {
	// Get config.
	cfg, err := getConfig()
	if err != nil {
		return err
	}
	utils.Info("config initialized")

	// Init newrelic.
	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.Newrelic.Name),
		newrelic.ConfigLicense(cfg.Newrelic.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		utils.Error(err.Error())
	} else {
		nrApp.WaitForConnection(10 * time.Second)
		defer nrApp.Shutdown(10 * time.Second)
		utils.AddLog(_nr.NewFromNewrelicApp(nrApp, _nr.ErrorLevel))
		utils.Info("newrelic initialized")
	}

	// Init db.
	db, err := newDB(cfg.DB)
	if err != nil {
		return err
	}
	utils.Info("database initialized")
	defer db.Client().Disconnect(context.Background())

	// Init discord.
	var discord discordRepository.Repository
	discord, err = discordClient.New(cfg.Discord.Token)
	if err != nil {
		return err
	}
	utils.Info("discord initialized")

	// Init user.
	var user userRepository.Repository = userDB.New(db)
	utils.Info("user repository initialized")

	// Init nyaa.
	var nyaa nyaaRepository.Repository = nyaaRSS.New()
	utils.Info("nyaa repository initialized")

	// Init template.
	var template templateRepository.Repository = templateClient.New(cfg.Discord.Prefix)
	utils.Info("template repository initialized")

	// Init service.
	service := service.New(discord, user, nyaa, template)
	utils.Info("service initialized")

	// Run cron.
	utils.Info("checking nyaa update...")
	if err := cron.New(service).Check(nrApp, cfg.Cron.Interval); err != nil {
		return err
	}

	utils.Info("done")
	return nil
}
