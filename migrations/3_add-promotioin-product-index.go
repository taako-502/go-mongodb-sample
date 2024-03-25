package migrations

import (
	"context"
	"os"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	migrate.Register(func(ctx context.Context, db *mongo.Database) error {
		if _, err := db.Collection(os.Getenv("PRODUCT_COLLECTION_NAME")).Indexes().CreateOne(context.TODO(), mongo.IndexModel{
			Keys:    bson.D{{Key: "promotion_expires_at", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0), // TTL index
		}); err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *mongo.Database) error {
		if _, err := db.Collection(os.Getenv("PRODUCT_COLLECTION_NAME")).Indexes().DropOne(context.TODO(), "promotion_expires_at_1"); err != nil {
			return err
		}
		return nil
	})
}
