package internal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName string = "nyaaXdiscord"
const collectionName string = "user"

// Database contains all function related with database.
type Database interface {
	GetUser(userID string) (User, bool, error)
	GetSubbedUser() ([]User, error)
	CreateUser(userID string, channelID string) error
	UpdateUser(user User) error
	Close() error
}

type database struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

// NewDB to create new database.
func NewDB(uri, user, pass string) (Database, error) {
	ctx := context.Background()

	// Start connection.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: user,
		Password: pass,
	}))
	if err != nil {
		return nil, err
	}

	// Ping test.
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &database{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
		ctx:        ctx,
	}, nil
}

// Close to close db connection.
func (d *database) Close() error {
	return d.client.Disconnect(d.ctx)
}

// GetUser to get user data.
func (d *database) GetUser(userID string) (user User, isExist bool, err error) {
	err = d.collection.FindOne(d.ctx, bson.M{"userId": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

// CreateUser to create new user.
func (d *database) CreateUser(userID string, channelID string) error {
	_, err := d.collection.InsertOne(d.ctx, User{
		UserID:    userID,
		ChannelID: channelID,
		Filter:    "0",
		Category:  "0_0",
		Queries:   []string{},
		Subscribe: false,
	})
	return err
}

// UpdateUser to update user data.
func (d *database) UpdateUser(user User) error {
	_, err := d.collection.UpdateOne(d.ctx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
	return err
}

// GetSubbedUser to get subscribed users.
func (d *database) GetSubbedUser() (users []User, err error) {
	cursor, err := d.collection.Find(d.ctx, bson.M{"subscribe": true})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(d.ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
