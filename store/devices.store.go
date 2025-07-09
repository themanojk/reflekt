package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/themanojk/reflekt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Insert(ctx context.Context, device *models.Device) (*models.Device, error)
	GetByMacAddress(ctx context.Context, macAddress string) (*models.Device, error)
}

type mongoStore struct {
	coll *mongo.Collection
}

// NewMongoStore returns a Store backed by the given Mongo client and database name.
func NewMongoStore(client *mongo.Client, dbName string) Store {
	coll := client.Database(dbName).Collection("devices")
	return &mongoStore{coll: coll}
}

func (m *mongoStore) Insert(ctx context.Context, device *models.Device) (*models.Device, error) {
	// Use a 5-second timeout per operation
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := m.coll.InsertOne(ctx, device)
	if err != nil {
		return nil, fmt.Errorf("failed to insert device: %w", err)
	}
	deviceId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}
	device.ID = deviceId

	return device, nil
}

func (m *mongoStore) GetByMacAddress(ctx context.Context, macAddress string) (*models.Device, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var device models.Device
	err := m.coll.FindOne(ctx, bson.M{"mac_address": macAddress}).Decode(&device)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("device not found")
	}
	if err != nil {
		return nil, err
	}
	return &device, nil
}
