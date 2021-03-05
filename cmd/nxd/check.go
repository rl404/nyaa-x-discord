package main

import (
	"log"
	"time"

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

	// Init discord.
	discord, err := internal.NewDiscord(cfg.Token)
	if err != nil {
		return err
	}

	// Init RSS.
	r := internal.NewRSS(db, discord, cfg.Interval, cfg.Location, logger)

	if err = r.Check(); err != nil {
		if logger != nil {
			if errLog := logger.Send("nxd-error", internal.LogError{
				Error:     err.Error(),
				CreatedAt: time.Now(),
			}); errLog != nil {
				log.Println(errLog)
			}
		}
	}

	return err
}
