package rss

import (
	"context"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/mmcdole/gofeed"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rl404/fairy/limit"
	"github.com/rl404/nyaa-x-discord/internal/domain/nyaa/entity"
	"github.com/rl404/nyaa-x-discord/internal/errors"
)

type rss struct {
	parser  *gofeed.Parser
	limiter limit.Limiter
	retrier *retrier.Retrier
}

// New to create new nyaa rss.
func New() *rss {
	limiter, _ := limit.New(limit.Atomic, 1, time.Second)

	parser := gofeed.NewParser()
	parser.Client = &http.Client{
		Timeout:   5 * time.Second,
		Transport: newrelic.NewRoundTripper(http.DefaultTransport),
	}

	return &rss{
		parser:  parser,
		limiter: limiter,
		retrier: retrier.New(retrier.ConstantBackoff(5, time.Second), nil),
	}
}

// GetFeeds to get nyaa feeds.
func (r *rss) GetFeeds(ctx context.Context, filter, category string, queries []string) ([]entity.Feed, error) {
	rawFeeds, err := r.getRawFeeds(ctx, filter, category, queries)
	if err != nil {
		return nil, errors.Wrap(ctx, err)
	}

	feeds := make([]entity.Feed, len(rawFeeds))
	for i, f := range rawFeeds {
		feeds[i] = entity.Feed{
			Title: f.Title,
			Link:  f.GUID,
			Size:  f.Extensions["nyaa"]["size"][0].Value,
			Date:  r.parseDate(f.Published),
		}
	}

	return feeds, nil
}

func (r *rss) getRawFeeds(ctx context.Context, filter, category string, queries []string) ([]*gofeed.Item, error) {
	var items []*gofeed.Item

	// Loop every 5 queries.
	for curr := 0; curr <= len(queries); curr += 5 {
		end := curr + 5
		if end > len(queries) {
			end = len(queries)
		}

		if curr == end {
			break
		}

		feed, err := r.parse(ctx, filter, category, queries[curr:end])
		if err != nil {
			return nil, errors.Wrap(ctx, err)
		}

		items = append(items, feed.Items...)
	}

	return items, nil
}

func (r *rss) parse(ctx context.Context, filter, category string, queries []string) (feed *gofeed.Feed, err error) {
	r.limiter.Take()

	if err2 := r.retrier.RunCtx(ctx, func(ctx context.Context) error {
		feed, err = r.parser.ParseURLWithContext(entity.GenerateURL(filter, category, queries, true), ctx)
		if err != nil {
			return err
		}
		return nil
	}); err2 != nil {
		return nil, errors.Wrap(ctx, err2)
	}

	return feed, nil
}

func (r *rss) parseDate(d string) time.Time {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	t, _ := time.Parse(layout, d)
	return t
}
