package client

import (
	"context"

	"github.com/rl404/nyaa-x-discord/internal/errors"
)

// CreateUserChannel to create user channel.
func (c *Client) CreateUserChannel(ctx context.Context, userID string) (string, error) {
	channel, err := c.session.UserChannelCreate(userID)
	if err != nil {
		return "", errors.Wrap(ctx, err)
	}
	return channel.ID, nil
}
