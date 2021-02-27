package main

import "github.com/rl404/nyaa-x-discord/internal"

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

	// Init discord.
	discord, err := internal.NewDiscord(cfg.Token)
	if err != nil {
		return err
	}

	// Init RSS.
	r := internal.NewRSS(db, discord, cfg.Interval, cfg.Location)

	return r.Check()
}
