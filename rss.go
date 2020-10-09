package main

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

// Feed is each feed data model.
type Feed struct {
	Title      string
	Link       string
	CategoryID string
	Category   string
	Size       string
	Date       time.Time
}

// fetchFeed to fetch feed from RSS.
func fetchFeed(URL string) (feeds *gofeed.Feed, err error) {
	return gofeed.NewParser().ParseURL(URL)
}

// getCleanFeed to clean feed model and get the recent feeds.
func getCleanFeed(user User) (feeds []Feed, err error) {
	f, err := fetchFeed(getNyaaQuery(user, true))
	if err != nil {
		return nil, err
	}

	limitDate := getLimitDate()

	for _, item := range f.Items {
		itemTime := parseDate(item.Published)

		// Skip old data.
		if itemTime.After(limitDate) {
			feeds = append(feeds, Feed{
				Title:      item.Title,
				Link:       item.GUID,
				CategoryID: item.Extensions["nyaa"]["categoryId"][0].Value,
				Category:   item.Extensions["nyaa"]["category"][0].Value,
				Size:       item.Extensions["nyaa"]["size"][0].Value,
				Date:       itemTime.In(cfg.Location),
			})
		}
	}

	fmt.Println(time.Now().Format("15:04:05"), "checking", user.UserID, len(feeds))
	return feeds, nil
}

// parseDate to parse string to time format.
func parseDate(d string) time.Time {
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	t, _ := time.Parse(layout, d)
	return t
}

// getLimitDate to get limit of recent data date.
func getLimitDate() time.Time {
	return time.Now().UTC().Add(time.Duration(-1*cfg.Interval) * time.Minute)
}
