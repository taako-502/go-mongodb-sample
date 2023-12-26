package example

import (
	"context"
	"fmt"
	customer_infrastructure "go-mongodb-sample/infrastructures/customers"
	order_infrastructure "go-mongodb-sample/infrastructures/orders"
	product_infrastructure "go-mongodb-sample/infrastructures/products"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Exammple(connectionString string, ctx context.Context, client *mongo.Client, dbname string) {
	// カスタマーを作成
	c := customer_infrastructure.NewCustomer(ctx,
		client.Database(dbname).Collection("customer"),
	)
	dto := customer_infrastructure.NewCustomerDTO("Alice", "alice@gmail.com", "Tokyo", nil)
	customer, err := c.Create(dto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", customer)

	// プロダクトを作成
	p := product_infrastructure.NewProduct(ctx,
		client.Database(dbname).Collection("product"),
	)
	productDto := product_infrastructure.NewProductDTO("Apple", "iPhone", 100000, 10, "Smartphone")
	product, err := p.Create(productDto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", product)

	// オーダーを作成
	o := order_infrastructure.NewOrderRepository(ctx,
		client.Database(dbname).Collection("order"),
	)
	orderDetailDto := order_infrastructure.NewOrderDetailDTO(product.ID, 100, 10000)
	orderDto := order_infrastructure.NewOrderDTO(customer.ID, []order_infrastructure.OrderDetailDTO{*orderDetailDto}, time.Now(), 100, "created")
	order, err := o.Create(orderDto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", order)

	// ドキュメントを取得するクエリ
	findedCustomer, err := c.FindOne(customer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found document: ", findedCustomer)
}
