package entity

// User is entity for user.
type User struct {
	ID        string
	UserID    string
	ChannelID string
	Filter    string
	Category  string
	Queries   []string
	Subscribe bool
}
