package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	nyaaEntity "github.com/rl404/nyaa-x-discord/internal/domain/nyaa/entity"
	"github.com/rl404/nyaa-x-discord/internal/domain/template/entity"
	"github.com/rl404/nyaa-x-discord/internal/errors"
	"github.com/rl404/nyaa-x-discord/internal/utils"
)

// HandlePing to handle ping.
func (s *service) HandlePing(ctx context.Context, m *discordgo.MessageCreate) error {
	_, err := s.discord.SendMessage(ctx, m.ChannelID, "pong")
	return errors.Wrap(ctx, err)
}

// HandleHelp to handle help.
func (s *service) HandleHelp(ctx context.Context, m *discordgo.MessageCreate) error {
	_, err := s.discord.SendMessageEmbed(ctx, m.ChannelID, s.template.GetHelp())
	return errors.Wrap(ctx, err)
}

// HandleFilter to handle filter command.
func (s *service) HandleFilter(ctx context.Context, user User, args []string) error {
	if len(args) == 1 {
		_, err := s.discord.SendMessageEmbed(ctx, user.ChannelID, s.template.GetFilters(entity.User{
			Filter:   user.Filter,
			Category: user.Category,
			Queries:  user.Queries,
		}))
		return errors.Wrap(ctx, err)
	}

	// Invalid command.
	if len(args) != 3 || args[1] != "set" || nyaaEntity.Filters.GetValueByKey(args[2]) == "" {
		return errors.Wrap(ctx, s.handleInvalid(ctx, user.ChannelID))
	}

	user.Filter = args[2]

	if err := s.user.UpdateFilterByUserID(ctx, user.UserID, user.Filter); err != nil {
		return errors.Wrap(ctx, err)
	}

	if err := s.handleSuccess(ctx, user); err != nil {
		return errors.Wrap(ctx, err)
	}

	return nil
}

// HandleCategory to handle category command.
func (s *service) HandleCategory(ctx context.Context, user User, args []string) error {
	if len(args) == 1 {
		_, err := s.discord.SendMessageEmbed(ctx, user.ChannelID, s.template.GetCategories(entity.User{
			Filter:   user.Filter,
			Category: user.Category,
			Queries:  user.Queries,
		}))
		return errors.Wrap(ctx, err)
	}

	// Invalid command.
	if len(args) != 3 || args[1] != "set" || nyaaEntity.Categories.GetValueByKey(args[2]) == "" {
		return errors.Wrap(ctx, s.handleInvalid(ctx, user.ChannelID))
	}

	user.Category = args[2]

	if err := s.user.UpdateCategoryByUserID(ctx, user.UserID, user.Category); err != nil {
		return errors.Wrap(ctx, err)
	}

	if err := s.handleSuccess(ctx, user); err != nil {
		return errors.Wrap(ctx, err)
	}

	return nil
}

// HandleQuery to handle query command.
func (s *service) HandleQuery(ctx context.Context, user User, args []string) error {
	if len(args) == 1 {
		_, err := s.discord.SendMessageEmbed(ctx, user.ChannelID, s.template.GetQueries(entity.User{
			Filter:   user.Filter,
			Category: user.Category,
			Queries:  user.Queries,
		}))
		return errors.Wrap(ctx, err)
	}

	// Invalid command.
	if len(args) < 3 {
		return errors.Wrap(ctx, s.handleInvalid(ctx, user.ChannelID))
	}

	switch args[1] {
	case "add":
		user.Queries = append(user.Queries, strings.Join(args[2:], " "))
	case "delete":
		var queries []string
		for i, q := range user.Queries {
			if !utils.InArray(args[2:], strconv.Itoa(i)) {
				queries = append(queries, q)
			}
		}
		user.Queries = queries
	default:
		return errors.Wrap(ctx, s.handleInvalid(ctx, user.ChannelID))
	}

	if err := s.user.UpdateQueriesByUserID(ctx, user.UserID, user.Queries); err != nil {
		return errors.Wrap(ctx, err)
	}

	if err := s.handleSuccess(ctx, user); err != nil {
		return errors.Wrap(ctx, err)
	}

	return nil
}

// HandleSubscribe to handle subscribe command.
func (s *service) HandleSubscribe(ctx context.Context, user User, args []string) error {
	if len(args) == 1 {
		_, err := s.discord.SendMessageEmbed(ctx, user.ChannelID, s.template.GetSubscribe(entity.User{
			Filter:    user.Filter,
			Category:  user.Category,
			Queries:   user.Queries,
			Subscribe: user.Subscribe,
		}))
		return errors.Wrap(ctx, err)
	}

	// Invalid command.
	if len(args) != 2 || (args[1] != "on" && args[1] != "off") {
		return errors.Wrap(ctx, s.handleInvalid(ctx, user.ChannelID))
	}

	user.Subscribe = args[1] == "on"

	if err := s.user.UpdateSubscribeByUserID(ctx, user.UserID, user.Subscribe); err != nil {
		return errors.Wrap(ctx, err)
	}

	if err := s.handleSuccess(ctx, user); err != nil {
		return errors.Wrap(ctx, err)
	}

	return nil
}

func (s *service) handleInvalid(ctx context.Context, channelID string) error {
	_, err := s.discord.SendMessage(ctx, channelID, s.template.GetInvalid())
	return errors.Wrap(ctx, err)
}

func (s *service) handleSuccess(ctx context.Context, user User) error {
	_, err := s.discord.SendMessageEmbed(ctx, user.ChannelID, s.template.GetUpdated(entity.User{
		Filter:   user.Filter,
		Category: user.Category,
		Queries:  user.Queries,
	}))
	return errors.Wrap(ctx, err)
}
