package client

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/rl404/nyaa-x-discord/internal/errors"
)

// SendMessage to send message.
func (c *Client) SendMessage(ctx context.Context, cID, content string) (string, error) {
	m, err := c.session.ChannelMessageSend(cID, content)
	if err != nil {
		return "", errors.Wrap(ctx, err)
	}
	return m.ID, nil
}

// SendMessageEmbed to send embed message.
func (c *Client) SendMessageEmbed(ctx context.Context, cID string, content *discordgo.MessageEmbed) (string, error) {
	m, err := c.session.ChannelMessageSendEmbed(cID, content)
	if err != nil {
		return "", errors.Wrap(ctx, err)
	}
	return m.ID, nil
}
