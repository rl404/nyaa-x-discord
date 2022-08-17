package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	_bot "github.com/rl404/nyaa-x-discord/internal/delivery/bot"
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

func bot() error {
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
	)
	if err != nil {
		utils.Error(err.Error())
	} else {
		defer nrApp.Shutdown(10 * time.Second)
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

	// Init bot.
	bot := _bot.New(service, cfg.Discord.Prefix)
	bot.RegisterHandler(nrApp)
	utils.Info("bot initialized")

	// Run bot.
	if err := bot.Run(); err != nil {
		return err
	}
	utils.Info("nxd is running...")
	defer bot.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	return nil
}
