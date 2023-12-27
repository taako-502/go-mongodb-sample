package main

import (
	"context"
	customer_controller "go-mongodb-sample/app/controllers/customer"
	order_controller "go-mongodb-sample/app/controllers/order"
	product_controller "go-mongodb-sample/app/controllers/product"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	// example.Exammple(connectionString, ctx, client, "testdb")

	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	customerController := customer_controller.NewCostumerController(ctx, client.Database("testdb").Collection("customer"))
	e.POST("/customer", customerController.Create)
	productController := product_controller.NewProductController(ctx, client.Database("testdb").Collection("product"))
	e.POST("/product", productController.Create)
	orderController := order_controller.NewOrderController(connectionString, "testdb", "order")
	e.GET("/orders/:customer_id", orderController.FindByCustomerID)
	e.POST("/order", orderController.Create)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
