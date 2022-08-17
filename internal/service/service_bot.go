package service

import (
	"github.com/bwmarrin/discordgo"
)

// Run to run discord bot.
func (s *service) Run() error {
	return s.discord.Run()
}

// Stop to stop discord bot.
func (s *service) Stop() error {
	return s.discord.Close()
}

// RegisterReadyHandler to register discord ready handler.
func (s *service) RegisterReadyHandler(fn func(*discordgo.Session, *discordgo.Ready)) {
	s.discord.AddReadyHandler(fn)
}

// RegisterMessageHandler to register discord message handler.
func (s *service) RegisterMessageHandler(fn func(*discordgo.Session, *discordgo.MessageCreate)) {
	s.discord.AddMessageHandler(fn)
}
