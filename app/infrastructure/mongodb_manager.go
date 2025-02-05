package infrastructure

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoDBManager struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoDBManager(ctx context.Context, connectionString string) (*mongoDBManager, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %w", err)
	}
	return &mongoDBManager{
		Client: client,
	}, nil
}
