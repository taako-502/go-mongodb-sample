package main

import (
	"context"
	"fmt"
	customer_infrastructure "go-mongodb-sample/infrastructures/customers"
	"log"
	"time"

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

	c := customer_infrastructure.NewCustomer(ctx, collection)
	dto := customer_infrastructure.NewCustomerDTO("Alice", "alice@gmail.com", "Tokyo", nil)
	customer, err := c.Create(dto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", customer)

	// ドキュメントを取得するクエリ
	findedCustomer, err := c.FindOne(customer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found document: ", findedCustomer)
}
