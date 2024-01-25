package migrations

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		opt := options.CreateCollection()
		err := db.CreateCollection(context.TODO(), "fuga", opt)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		err := db.Collection("fuga").Drop(context.TODO())
		if err != nil {
			return err
		}
		return nil
	})
}
