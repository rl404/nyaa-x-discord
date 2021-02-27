package internal

import "github.com/bwmarrin/discordgo"

// Discord contains all basic discord function.
type Discord interface {
	AddHandler(func(*discordgo.Session, *discordgo.MessageCreate))
	Run() error
	SendMessage(channelID string, message string) error
	SendEmbedMessage(channelID string, message discordgo.MessageEmbed) error
	Close() error
}

type discord struct {
	client *discordgo.Session
}

// NewDiscord to create new discord client.
func NewDiscord(token string) (Discord, error) {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &discord{
		client: client,
	}, nil
}

// AddHandler to add message handler.
func (d *discord) AddHandler(handler func(*discordgo.Session, *discordgo.MessageCreate)) {
	d.client.AddHandler(handler)
}

// SendMessage to send message to channel.
func (d *discord) SendMessage(chID string, content string) error {
	_, err := d.client.ChannelMessageSend(chID, content)
	return err
}

// SendEmbedMessage to send embed message to channel.
func (d *discord) SendEmbedMessage(chID string, content discordgo.MessageEmbed) error {
	_, err := d.client.ChannelMessageSendEmbed(chID, &content)
	return err
}

// Run to login and start discord bot.
func (d *discord) Run() error {
	return d.client.Open()
}

// Close to stop discord bot.
func (d *discord) Close() error {
	return d.client.Close()
}
