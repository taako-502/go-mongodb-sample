package main

import (
	"context"
	"fmt"
	customer_infrastructure "go-mongodb-sample/infrastructures/customers"
	product_infrastructure "go-mongodb-sample/infrastructures/products"
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

	// カスタマーを作成
	c := customer_infrastructure.NewCustomer(ctx,
		client.Database("testdb").Collection("customer"),
	)
	dto := customer_infrastructure.NewCustomerDTO("Alice", "alice@gmail.com", "Tokyo", nil)
	customer, err := c.Create(dto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", customer)

	// プロダクトを作成
	p := product_infrastructure.NewProduct(ctx,
		client.Database("testdb").Collection("product"),
	)
	productDto := product_infrastructure.NewProductDTO("Apple", "iPhone", 100000, 10, "Smartphone")
	product, err := p.Create(productDto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", product)

	// ドキュメントを取得するクエリ
	findedCustomer, err := c.FindOne(customer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found document: ", findedCustomer)
}
