package client

import (
	"context"

	"github.com/rl404/fairy/errors/stack"
)

// CreateUserChannel to create user channel.
func (c *Client) CreateUserChannel(ctx context.Context, userID string) (string, error) {
	channel, err := c.session.UserChannelCreate(userID)
	if err != nil {
		return "", stack.Wrap(ctx, err)
	}
	return channel.ID, nil
}
