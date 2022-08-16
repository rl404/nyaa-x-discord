package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/eapache/go-resiliency/retrier"
	"github.com/mmcdole/gofeed"
)

// RSS contains function to check RSS.
type RSS interface {
	Check() error
}

type rss struct {
	db       Database
	discord  Discord
	interval int
}

// Feed is each feed data model.
type Feed struct {
	Title      string
	Link       string
	CategoryID string
	Category   string
	Size       string
	Date       time.Time
}

// NewRSS to create new RSS.
func NewRSS(db Database, d Discord, interval int) RSS {
	return &rss{
		db:       db,
		discord:  d,
		interval: interval,
	}
}

// Check to check user's new feeds.
func (r *rss) Check() error {
	users, err := r.db.GetSubbedUser()
	if err != nil {
		return err
	}

	// Check every user.
	for _, user := range users {
		feeds, err := r.getFeeds(user)
		if err != nil {
			continue
		}

		if len(feeds) > 0 {
			var titles []string
			for _, f := range feeds {
				titles = append(titles, f.Title)
			}
		}

		log.Println(user.UserID, len(feeds))

		if len(feeds) > 0 {
			// Send message if there are new feeds.
			err = r.sendFeed(feeds, user)
		}
	}

	return nil
}

func (r *rss) getFeeds(user User) (feeds []Feed, err error) {
	items, err := r.getRawFeeds(user)
	if err != nil {
		return nil, err
	}

	limitDate := r.getLimitDate()

	for _, item := range items {
		itemTime := r.parseDate(item.Published)

		// Skip old data.
		if itemTime.After(limitDate) {
			feeds = append(feeds, Feed{
				Title:      item.Title,
				Link:       item.GUID,
				CategoryID: item.Extensions["nyaa"]["categoryId"][0].Value,
				Category:   item.Extensions["nyaa"]["category"][0].Value,
				Size:       item.Extensions["nyaa"]["size"][0].Value,
			})
		}
	}

	return feeds, nil
}

func (r *rss) getRawFeeds(user User) (feeds []*gofeed.Item, err error) {
	// If no query, just parse without query.
	if len(user.Queries) == 0 {
		f, err := r.parse(user)
		if err != nil {
			return nil, err
		}
		return f.Items, nil
	}

	// Loop every 5 queries.
	for curr := 0; curr < len(user.Queries); curr += 5 {
		end := curr + 5
		if end > len(user.Queries) {
			end = len(user.Queries)
		}

		tmp := user
		tmp.Queries = user.Queries[curr:end]
		f, err := r.parse(tmp)
		if err != nil {
			return nil, err
		}

		feeds = append(feeds, f.Items...)
	}

	return feeds, nil
}

func (r *rss) parse(user User) (f *gofeed.Feed, err error) {
	rt := retrier.New(retrier.ConstantBackoff(5, time.Second), nil)
	err2 := rt.Run(func() error {
		f, err = gofeed.NewParser().ParseURL(getNyaaQuery(user, true))
		if err != nil {
			return err
		}
		return nil
	})
	return f, err2
}

func (r *rss) getLimitDate() time.Time {
	return time.Now().UTC().Add(time.Duration(-1*r.interval) * time.Minute)
}

func (r *rss) parseDate(d string) time.Time {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	t, _ := time.Parse(layout, d)
	return t
}

func (r *rss) sendFeed(feeds []Feed, user User) error {
	msg := discordgo.MessageEmbed{
		Color: blueColor,
	}

	// Discord limits 25 fields per message.
	for curr := 0; curr < len(feeds); curr += 25 {
		var fields []*discordgo.MessageEmbedField

		end := curr + 25
		if end > len(feeds) {
			end = len(feeds)
		}

		for _, feed := range feeds[curr:end] {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name: ellipsis(feed.Title, 100),
				Value: fmt.Sprintf("[link](%s)  •  `%s`",
					feed.Link,
					feed.Size),
			})
		}

		msg.Fields = fields

		if err := r.discord.SendEmbedMessage(user.ChannelID, msg); err != nil {
			return err
		}
	}

	return nil
}
