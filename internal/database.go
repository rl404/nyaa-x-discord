package internal

import (
	"context"
	"time"

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
}

// NewDB to create new database.
func NewDB(uri, user, pass string) (Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start connection.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: user,
		Password: pass,
	}))
	if err != nil {
		return nil, err
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	// Ping test.
	if err = client.Ping(ctx2, nil); err != nil {
		return nil, err
	}

	return &database{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}, nil
}

// Close to close db connection.
func (d *database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return d.client.Disconnect(ctx)
}

// GetUser to get user data.
func (d *database) GetUser(userID string) (user User, isExist bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = d.collection.FindOne(ctx, bson.M{"userId": userID}).Decode(&user)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.collection.InsertOne(ctx, User{
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := d.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.D{{Key: "$set", Value: user}})
	return err
}

// GetSubbedUser to get subscribed users.
func (d *database) GetSubbedUser() (users []User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := d.collection.Find(ctx, bson.M{"subscribe": true})
	if err != nil {
		return nil, err
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	if err = cursor.All(ctx2, &users); err != nil {
		return nil, err
	}
	return users, nil
}
