package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rl404/nyaa-x-discord/internal"
)

func bot() error {
	// Get config.
	cfg, err := internal.GetConfig()
	if err != nil {
		return err
	}

	// Init db.
	db, err := internal.NewDB(cfg.DB.URI, cfg.DB.User, cfg.DB.Password)
	if err != nil {
		return err
	}
	defer db.Close()

	// Init discord.
	discord, err := internal.NewDiscord(cfg.Token)
	if err != nil {
		return err
	}

	// Init handler.
	h := internal.NewHandler(db, cfg.Prefix)

	// Add message handler.
	discord.AddHandler(h.Handler())

	// Run bot.
	if err = discord.Run(); err != nil {
		return err
	}
	defer discord.Close()

	log.Println("discord bot is running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit

	return nil
}
