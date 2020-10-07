package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Read and set config.
	err := setConfig()
	if err != nil {
		fmt.Println("config -", err)
		return
	}

	// Init discord client.
	err = initDiscord()
	defer discord.Close()
	if err != nil {
		fmt.Println("discord -", err)
		return
	}

	// Start Nyaa RSS checker and listening discord bot.
	startScheduler()
}

// startScheduler to start running scheduler.
func startScheduler() {
	// Add job.
	worker := newScheduler()
	worker.add(context.Background(), feedJob, time.Duration(cfg.Interval)*time.Minute)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit

	// Clean everything.
	worker.stop()
	discord.Close()
}

// feedJob is scheduler job to notify each user.
func feedJob(ctx context.Context) {
	users, err := getSubsUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println(time.Now().Format("15:03:04"), "checking ", user.UserID)
		checkFeed(user)

		// Sleep to prevent spamming.
		time.Sleep(time.Second)
	}
}

// checkFeed to check feeds update and send if there is new update.
func checkFeed(user User) {
	feeds, err := getCleanFeed(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(feeds) > 0 {
		sendFeed(feeds, user)
	}
}
