package client

import "github.com/bwmarrin/discordgo"

// Client is discord client.
type Client struct {
	session *discordgo.Session
}

// New to create new discord client.
func New(token string) (*Client, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &Client{
		session: session,
	}, nil
}

// Run to login and start discord bot.
func (c *Client) Run() error {
	return c.session.Open()
}

// Close to stop discord bot.
func (c *Client) Close() error {
	return c.session.Close()
}

// AddReadyHandler to add ready handler.
func (c *Client) AddReadyHandler(handler func(*discordgo.Session, *discordgo.Ready)) {
	c.session.AddHandler(handler)
}

// AddMessageHandler to add message handler.
func (c *Client) AddMessageHandler(handler func(*discordgo.Session, *discordgo.MessageCreate)) {
	c.session.AddHandler(handler)
}
