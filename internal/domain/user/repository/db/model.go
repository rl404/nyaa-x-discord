package db

import (
	"github.com/rl404/nyaa-x-discord/internal/domain/user/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"userId"`
	ChannelID string             `bson:"channelId"`
	Filter    string             `bson:"filter"`
	Category  string             `bson:"category"`
	Queries   []string           `bson:"queries"`
	Subscribe bool               `bson:"subscribe"`
}

func (u *user) toEntity() *entity.User {
	return &entity.User{
		ID:        u.ID.String(),
		UserID:    u.UserID,
		ChannelID: u.ChannelID,
		Filter:    u.Filter,
		Category:  u.Category,
		Queries:   u.Queries,
		Subscribe: u.Subscribe,
	}
}

func (db *db) toEntities(data []user) []*entity.User {
	users := make([]*entity.User, len(data))
	for i, u := range data {
		users[i] = u.toEntity()
	}
	return users
}
