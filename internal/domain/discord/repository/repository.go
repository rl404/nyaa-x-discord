package repository

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

// Repository contains functions for discord domain.
type Repository interface {
	Run() error
	Close() error

	AddReadyHandler(func(*discordgo.Session, *discordgo.Ready))
	AddMessageHandler(func(*discordgo.Session, *discordgo.MessageCreate))

	SendMessage(ctx context.Context, channelID, content string) (string, error)
	SendMessageEmbed(ctx context.Context, channelID string, content *discordgo.MessageEmbed) (string, error)

	CreateUserChannel(ctx context.Context, userID string) (string, error)
}
