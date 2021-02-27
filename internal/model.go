package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

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
