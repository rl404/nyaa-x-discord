package entity

// User is entity for user.
type User struct {
	Filter    string
	Category  string
	Queries   []string
	Subscribe bool
}

// Feed is entity for rss feed.
type Feed struct {
	Title string
	Link  string
	Size  string
}
