package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/taako-502/go-mongodb-sample/app/infrastructure/customer_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/order_infrastructure"
	"github.com/taako-502/go-mongodb-sample/app/infrastructure/product_infrastructure"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Exammple(connectionString string, ctx context.Context, dbname string) {
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// カスタマーを作成
	c := customer_infrastructure.NewCustomerRepository(ctx, client.Database(dbname))
	dto := customer_infrastructure.NewCustomerDTO("Alice", "alice@gmail.com", "Tokyo")
	customer, err := c.Create(dto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", customer)

	// プロダクトを作成
	p := product_infrastructure.NewProductRepository(ctx, client.Database(dbname))
	productDto := product_infrastructure.NewProductDTO("Apple", "iPhone", 100000, 10, "Smartphone")
	product, err := p.Create(productDto)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", product)

	// オーダーを作成
	o := order_infrastructure.NewOrderRepository(ctx, client.Database(dbname))
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
