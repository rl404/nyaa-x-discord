package internal

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is model for each user.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	ChannelID string             `bson:"channelId"`
	Filter    string             `bson:"filter"`
	Category  string             `bson:"category"`
	Queries   []string           `bson:"queries"`
	Subscribe bool               `bson:"subscribe"`
}

// LogData is model for logging.
type LogData struct {
	UserID    string    `json:"user_id"`
	Titles    []string  `json:"titles"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

// LogError is model for error logging.
type LogError struct {
	Error     string    `json:"error"`
	CreatedAt time.Time `json:"created_at"`
}
