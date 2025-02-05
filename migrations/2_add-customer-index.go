package migrations

import (
	"context"
	"os"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func init() {
	migrate.Register(func(ctx context.Context, db *mongo.Database) error {
		customerIndexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}}, // 1は昇順を意味する
			Options: options.Index().SetUnique(true),
		}
		if _, err := db.Collection(os.Getenv("CUSTOMER_COLLECTION_NAME")).Indexes().CreateOne(context.TODO(), customerIndexModel); err != nil {
			return err
		}
		return nil
	}, func(ctx context.Context, db *mongo.Database) error {
		if _, err := db.Collection(os.Getenv("CUSTOMER_COLLECTION_NAME")).Indexes().DropOne(context.TODO(), "email_1"); err != nil {
			return err
		}
		return nil
	})
}
