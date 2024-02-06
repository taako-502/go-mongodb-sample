package infrastructure

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBManager struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoDBManager(ctx context.Context, connectionString string) (*mongoDBManager, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, errors.Wrap(err, "mongo.Connect")
	}
	return &mongoDBManager{
		Client: client,
	}, nil
}
