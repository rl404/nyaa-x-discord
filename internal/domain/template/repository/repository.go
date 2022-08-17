package repository

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rl404/nyaa-x-discord/internal/domain/template/entity"
)

// Repository contains functions for template domain.
type Repository interface {
	GetFirstTime(username string) *discordgo.MessageEmbed
	GetHelp() *discordgo.MessageEmbed
	GetFilters(user entity.User) *discordgo.MessageEmbed
	GetCategories(user entity.User) *discordgo.MessageEmbed
	GetQueries(user entity.User) *discordgo.MessageEmbed
	GetSubscribe(user entity.User) *discordgo.MessageEmbed
	GetInvalid() string
	GetUpdated(user entity.User) *discordgo.MessageEmbed
	GetNyaaUpdate(feeds []entity.Feed) *discordgo.MessageEmbed
}
