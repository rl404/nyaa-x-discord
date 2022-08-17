package entity

import "time"

// Feed is entity for rss feed.
type Feed struct {
	Title string
	Link  string
	Size  string
	Date  time.Time
}
