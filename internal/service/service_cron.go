package service

import (
	"context"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/nyaa-x-discord/internal/domain/template/entity"
	"github.com/rl404/nyaa-x-discord/internal/errors"
)

// CheckNyaa to check nyaa update.
func (s *service) CheckNyaa(ctx context.Context, nrApp *newrelic.Application, interval time.Duration) error {
	users, err := s.user.GetSubscribedUsers(ctx)
	if err != nil {
		return errors.Wrap(ctx, err)
	}

	limitDate := time.Now().UTC().Add(-interval)

	for _, user := range users {
		feeds, err := s.nyaa.GetFeeds(ctx, user.Filter, user.Category, user.Queries)
		if err != nil {
			errors.Wrap(ctx, err)
			continue
		}

		var templateFeeds []entity.Feed
		for _, f := range feeds {
			// Skip old data.
			if f.Date.Before(limitDate) {
				continue
			}

			templateFeeds = append(templateFeeds, entity.Feed{
				Title: f.Title,
				Link:  f.Link,
				Size:  f.Size,
			})
		}

		// No new updates.
		if len(templateFeeds) == 0 {
			continue
		}

		nrApp.RecordCustomEvent("NyaaUpdate", map[string]interface{}{
			"user_id": user.UserID,
			"count":   len(templateFeeds),
		})

		// Discord limit 25 fields per message.
		for curr := 0; curr < len(templateFeeds); curr += 25 {
			end := curr + 25
			if end > len(feeds) {
				end = len(feeds)
			}

			// Notify user.
			if _, err := s.discord.SendMessageEmbed(ctx, user.ChannelID, s.template.GetNyaaUpdate(templateFeeds[curr:end])); err != nil {
				errors.Wrap(ctx, err)
			}
		}
	}

	return nil
}
