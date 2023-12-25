package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDBサーバーへの接続文字列
	connectionString := "mongodb://localhost:27017"

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// データベースとコレクションを選択
	collection := client.Database("testdb").Collection("testcollection")

	// 挿入するドキュメント
	document := bson.D{{Key: "name", Value: "Alice"}, {Key: "age", Value: 25}}

	// ドキュメントを挿入
	insertResult, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)

	// ドキュメントを取得するクエリ
	var result bson.D
	err = collection.FindOne(ctx, bson.D{{Key: "name", Value: "Alice"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found document: ", result)
}
