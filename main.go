package main

import (
	"context"
	"go-mongodb-sample/internal/example"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBサーバーへの接続文字列
const connectionString = "mongodb://localhost:27017"

func main() {
	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// サンプルを実行
	example.Exammple(connectionString, ctx, client)
}
