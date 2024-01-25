package migrations

import (
	"context"
	"fmt"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	fmt.Println("init")
	migrate.Register(func(db *mongo.Database) error {
		opt := options.CreateCollection()
		if err := db.CreateCollection(context.TODO(), "fuga", opt); err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		if err := db.Collection("fuga").Drop(context.TODO()); err != nil {
			return err
		}
		return nil
	})
}
