package bot

import (
	"context"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/nyaa-x-discord/internal/errors"
)

func (b *Bot) messageHandler(nrApp *newrelic.Application) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		ctx := errors.Init(context.Background())
		defer b.log(ctx)

		// Ignore all messages created by the bot itself.
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Command and prefix check.
		if b.prefixCheck(m.Content) {
			return
		}

		// Remove prefix.
		m.Content = b.cleanPrefix(m.Content)

		// Get user data.
		user, err := b.service.GetUserByUserID(ctx, m.Author.ID)
		if err != nil {
			errors.Wrap(ctx, err)
			return
		}

		// Handle first time.
		if user == nil {
			if err := b.service.HandleFirstTime(ctx, m); err != nil {
				errors.Wrap(ctx, err)
				return
			}
		}

		// Get arguments.
		r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
		args := r.FindAllString(m.Content, -1)

		tx := nrApp.StartTransaction("Command " + args[0])
		defer tx.End()

		ctx = newrelic.NewContext(ctx, tx)

		switch args[0] {
		case "ping":
			errors.Wrap(ctx, b.service.HandlePing(ctx, m))
		case "help":
			errors.Wrap(ctx, b.service.HandleHelp(ctx, m))
		case "filter":
			errors.Wrap(ctx, b.service.HandleFilter(ctx, *user, args))
		case "category":
			errors.Wrap(ctx, b.service.HandleCategory(ctx, *user, args))
		case "query":
			errors.Wrap(ctx, b.service.HandleQuery(ctx, *user, args))
		case "subscribe":
			errors.Wrap(ctx, b.service.HandleSubscribe(ctx, *user, args))
		}
	}
}

func (b *Bot) prefixCheck(cmd string) bool {
	return len(cmd) <= len(b.prefix) || cmd[:len(b.prefix)] != b.prefix
}

func (b *Bot) cleanPrefix(cmd string) string {
	return strings.TrimSpace(cmd[len(b.prefix):])
}
