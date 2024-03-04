package test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cIDatabaseConfig struct {
	database string
	uRL      string
	Client   *mongo.Client
}

func NewCIDatabaseConfig(ctx context.Context) (*cIDatabaseConfig, error) {
	if os.Getenv("ENV") != "ci" {
		// 実行中のファイルのディレクトリを取得
		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)

		// basepath から .env ファイルの絶対パスを組み立てる
		envPath := filepath.Join(basepath, "../.env")

		if err := godotenv.Load(envPath); err != nil {
			return nil, fmt.Errorf("godotenv.Load(%s): %w", envPath, err)
		}
	}

	db := cIDatabaseConfig{
		database: os.Getenv("MONGO_CI_DATABASE_NAME"),
		uRL:      os.Getenv("MONGO_CI_CONNECTION_STRING"),
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.uRL))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect(ctx, options.Client().ApplyURI(%s): %w", db.uRL, err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err = client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("client.Ping(pingCtx, nil): %w", err)
	}

	db.Client = client
	return &db, nil
}

func (db cIDatabaseConfig) GetDatabaseInstance() *mongo.Database {
	return db.Client.Database(db.database)
}
