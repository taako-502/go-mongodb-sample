package migrations

import (
	"context"
	"fmt"
	"log"
	"time"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(ctx context.Context, host, user, password, database string) (*mongo.Database, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017", user, password, host)
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	// 接続の確認
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(database)

	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
