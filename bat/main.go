package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("環境変数の読込に失敗しました: %v\r\n", err)
	}

	mongodb := os.Getenv("MONGO_CONNECTION_STRING")

	// MongoDBサーバーへの接続設定
	clientOptions := options.Client().ApplyURI(mongodb)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// データベースとコレクションの選択
	collection := client.Database("testdb").Collection("hoge")

	// クエリの実行
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// 結果の取得
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result["_id"])
	}
}
