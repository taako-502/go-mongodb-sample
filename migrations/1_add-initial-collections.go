package migrations

import (
	"context"
	"os"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	migrate.Register(func(ctx context.Context, db *mongo.Database) error {
		opt := options.CreateCollection()
		if err := db.CreateCollection(context.TODO(), os.Getenv("CUSTOMER_COLLECTION_NAME"), opt); err != nil {
			return err
		}
		if err := db.CreateCollection(context.TODO(), os.Getenv("ORDER_COLLECTION_NAME"), opt); err != nil {
			return err
		}
		if err := db.CreateCollection(context.TODO(), os.Getenv("PRODUCT_COLLECTION_NAME"), opt); err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *mongo.Database) error {
		if err := db.Collection(os.Getenv("CUSTOMER_COLLECTION_NAME")).Drop(context.TODO()); err != nil {
			return err
		}
		if err := db.Collection(os.Getenv("ORDER_COLLECTION_NAME")).Drop(context.TODO()); err != nil {
			return err
		}
		if err := db.Collection("PRODUCT_COLLECTION_NAME").Drop(context.TODO()); err != nil {
			return err
		}
		return nil
	})
}
