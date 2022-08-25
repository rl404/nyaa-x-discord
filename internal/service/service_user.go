package service

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/nyaa-x-discord/internal/errors"
)

// User is user model.
type User struct {
	UserID    string
	ChannelID string
	Filter    string
	Category  string
	Queries   []string
	Subscribe bool
}

// GetUserByUserID to get user by user id.
func (s *service) GetUserByUserID(ctx context.Context, UserID string) (*User, error) {
	user, err := s.user.GetByUserID(ctx, UserID)
	if err != nil {
		return nil, errors.Wrap(ctx, err)
	}
	if user == nil {
		return nil, nil
	}
	return &User{
		UserID:    user.UserID,
		ChannelID: user.ChannelID,
		Filter:    user.Filter,
		Category:  user.Category,
		Queries:   user.Queries,
		Subscribe: user.Subscribe,
	}, nil
}

// HandleFirstTime to handle first time user.
func (s *service) HandleFirstTime(ctx context.Context, m *discordgo.MessageCreate) error {
	// Get DM channel.
	channelID, err := s.discord.CreateUserChannel(ctx, m.Author.ID)
	if err != nil {
		return errors.Wrap(ctx, err)
	}

	// Create new user.
	if err = s.user.Create(ctx, m.Author.ID, channelID); err != nil {
		return errors.Wrap(ctx, err)
	}

	_, err = s.discord.SendMessageEmbed(ctx, channelID, s.template.GetFirstTime(m.Author.Username))
	return errors.Wrap(ctx, err)
}
