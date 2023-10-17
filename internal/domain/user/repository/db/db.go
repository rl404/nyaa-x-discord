package db

import (
	"context"

	"github.com/rl404/fairy/errors/stack"
	"github.com/rl404/nyaa-x-discord/internal/domain/user/entity"
	"github.com/rl404/nyaa-x-discord/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	db *mongo.Collection
}

// New to create new database.
func New(mdb *mongo.Database) *db {
	return &db{
		db: mdb.Collection("user"),
	}
}

// GetByUserID to get user data.
func (db *db) GetByUserID(ctx context.Context, userID string) (*entity.User, error) {
	var user user
	if err := db.db.FindOne(ctx, bson.M{"userId": userID}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, stack.Wrap(ctx, err, errors.ErrInternalDB)
	}
	return user.toEntity(), nil
}

// Create to create new user.
func (db *db) Create(ctx context.Context, userID, channelID string) error {
	if _, err := db.db.InsertOne(ctx, user{
		UserID:    userID,
		ChannelID: channelID,
		Filter:    "0",
		Category:  "0_0",
		Queries:   []string{},
		Subscribe: false,
	}); err != nil {
		return stack.Wrap(ctx, err, errors.ErrInternalDB)
	}
	return nil
}

// UpdateFilterByUserID to update filter user data.
func (db *db) UpdateFilterByUserID(ctx context.Context, userID, filter string) error {
	if _, err := db.db.UpdateOne(ctx, bson.D{{Key: "userId", Value: userID}}, bson.D{{Key: "$set", Value: bson.D{{Key: "filter", Value: filter}}}}); err != nil {
		return stack.Wrap(ctx, err, errors.ErrInternalDB)
	}
	return nil
}

// UpdateCategoryByUserID to update category user data.
func (db *db) UpdateCategoryByUserID(ctx context.Context, userID, category string) error {
	if _, err := db.db.UpdateOne(ctx, bson.D{{Key: "userId", Value: userID}}, bson.D{{Key: "$set", Value: bson.D{{Key: "category", Value: category}}}}); err != nil {
		return stack.Wrap(ctx, err, errors.ErrInternalDB)
	}
	return nil
}

// UpdateQueriesByUserID to update queries user data.
func (db *db) UpdateQueriesByUserID(ctx context.Context, userID string, queries []string) error {
	if _, err := db.db.UpdateOne(ctx, bson.D{{Key: "userId", Value: userID}}, bson.D{{Key: "$set", Value: bson.D{{Key: "queries", Value: queries}}}}); err != nil {
		return stack.Wrap(ctx, err, errors.ErrInternalDB)
	}
	return nil
}

// UpdateSubscribeByUserID to update subscribe user data.
func (db *db) UpdateSubscribeByUserID(ctx context.Context, userID string, subscribe bool) error {
	if _, err := db.db.UpdateOne(ctx, bson.D{{Key: "userId", Value: userID}}, bson.D{{Key: "$set", Value: bson.D{{Key: "subscribe", Value: subscribe}}}}); err != nil {
		return stack.Wrap(ctx, err, errors.ErrInternalDB)
	}
	return nil
}

// GetSubscribedUsers to get subscribed users.
func (db *db) GetSubscribedUsers(ctx context.Context) ([]*entity.User, error) {
	cursor, err := db.db.Find(ctx, bson.D{{Key: "subscribe", Value: true}})
	if err != nil {
		return nil, stack.Wrap(ctx, err, errors.ErrInternalDB)
	}

	var users []user
	if err := cursor.All(ctx, &users); err != nil {
		return nil, stack.Wrap(ctx, err, errors.ErrInternalDB)
	}

	return db.toEntities(users), nil
}
