package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database client.
var dbClient *mongo.Client
var dbCollection *mongo.Collection
var dbCtx = context.TODO()

// Database and collection name.
const dbName string = "nyaaXdiscord"
const collectionName string = "user"

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

// setConnection to prepare db connection.
func setConnection() error {
	// Start connection.
	dbClient, err := mongo.Connect(dbCtx, options.Client().ApplyURI(cfg.DB))
	if err != nil {
		return err
	}

	// Ping test.
	err = dbClient.Ping(dbCtx, nil)
	if err != nil {
		return err
	}

	// Set database and collection name.
	dbCollection = dbClient.Database(dbName).Collection(collectionName)

	return nil
}

// getUser to get user data.
func getUser(userID string) (user User, isExist bool, err error) {
	err = dbCollection.FindOne(dbCtx, bson.M{"userId": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

// createUser to create new query for user.
func createUser(userID string, channelID string) error {
	_, err := dbCollection.InsertOne(dbCtx, User{
		UserID:    userID,
		ChannelID: channelID,
		Filter:    "0",
		Category:  "0_0",
		Queries:   []string{},
		Subscribe: false,
	})
	return err
}

// getSubsUser to get subscribed users.
func getSubsUser() (users []User, err error) {
	cursor, err := dbCollection.Find(dbCtx, bson.M{"subscribe": true})
	if err != nil {
		return nil, err
	}

	err = cursor.All(dbCtx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
