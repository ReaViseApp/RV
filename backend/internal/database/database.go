package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewDatabase(uri, dbName string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	database := client.Database(dbName)

	return &Database{
		Client:   client,
		Database: database,
	}, nil
}

func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}

// Collection helpers
func (db *Database) Users() *mongo.Collection {
	return db.Database.Collection("users")
}

func (db *Database) Posts() *mongo.Collection {
	return db.Database.Collection("posts")
}

func (db *Database) Comments() *mongo.Collection {
	return db.Database.Collection("comments")
}

func (db *Database) Messages() *mongo.Collection {
	return db.Database.Collection("messages")
}

func (db *Database) Transactions() *mongo.Collection {
	return db.Database.Collection("transactions")
}

func (db *Database) NFTListings() *mongo.Collection {
	return db.Database.Collection("nft_listings")
}

func (db *Database) Likes() *mongo.Collection {
	return db.Database.Collection("likes")
}

func (db *Database) Follows() *mongo.Collection {
	return db.Database.Collection("follows")
}
