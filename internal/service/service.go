package service

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/newrelic/go-agent/v3/newrelic"
	discordRepository "github.com/rl404/nyaa-x-discord/internal/domain/discord/repository"
	nyaaRepository "github.com/rl404/nyaa-x-discord/internal/domain/nyaa/repository"
	templateRepository "github.com/rl404/nyaa-x-discord/internal/domain/template/repository"
	userRepository "github.com/rl404/nyaa-x-discord/internal/domain/user/repository"
)

// Service contains functions for service.
type Service interface {
	Run() error
	Stop() error

	RegisterReadyHandler(func(*discordgo.Session, *discordgo.Ready))
	RegisterMessageHandler(func(*discordgo.Session, *discordgo.MessageCreate))

	HandlePing(ctx context.Context, m *discordgo.MessageCreate) error
	HandleHelp(ctx context.Context, m *discordgo.MessageCreate) error
	HandleFirstTime(ctx context.Context, m *discordgo.MessageCreate) error
	HandleFilter(ctx context.Context, user User, args []string) error
	HandleCategory(ctx context.Context, user User, args []string) error
	HandleQuery(ctx context.Context, user User, args []string) error
	HandleSubscribe(ctx context.Context, user User, args []string) error

	GetUserByUserID(ctx context.Context, userID string) (*User, error)

	CheckNyaa(ctx context.Context, nrApp *newrelic.Application, interval time.Duration) error
}

type service struct {
	discord  discordRepository.Repository
	user     userRepository.Repository
	nyaa     nyaaRepository.Repository
	template templateRepository.Repository
}

// New to create new service.
func New(
	discord discordRepository.Repository,
	user userRepository.Repository,
	nyaa nyaaRepository.Repository,
	template templateRepository.Repository,
) Service {
	return &service{
		discord:  discord,
		user:     user,
		nyaa:     nyaa,
		template: template,
	}
}
