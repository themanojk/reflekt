package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoClient connects to Mongo using the given URI and returns the client.
// Caller should call client.Disconnect when done.
func NewMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	// Use a 10-second timeout for the initial connection
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}
