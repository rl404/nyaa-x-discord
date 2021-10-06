package main

import (
	"github.com/rl404/nyaa-x-discord/internal"
)

func check() error {
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

	// Init elasticsearch.
	var logger internal.Logger
	if len(cfg.ES.Address) > 0 {
		if logger, err = internal.NewES(cfg.ES.Address, cfg.ES.User, cfg.ES.Password); err != nil {
			return err
		}
	}
	return nil
	// Init discord.
	discord, err := internal.NewDiscord(cfg.Token)
	if err != nil {
		return err
	}

	// Init RSS.
	r := internal.NewRSS(db, discord, cfg.Interval, logger)

	// Run check.
	err = r.Check()

	internal.HandleError(logger, err)

	return err
}
