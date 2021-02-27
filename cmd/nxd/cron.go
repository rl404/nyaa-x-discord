package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rl404/nyaa-x-discord/internal"
)

// Adding cron to docker is kinda weird, so let's create
// our own cron then.
func cron() error {
	// Get config.
	cfg, err := internal.GetConfig()
	if err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	go func() {
		for {
			if err := check(); err != nil {
				log.Println(err)
			}
			time.Sleep(time.Duration(cfg.Interval) * time.Minute)
		}
	}()

	<-quit

	return nil
}
