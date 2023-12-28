package main

import (
	"context"
	customer_controller "go-mongodb-sample/app/controllers/customer"
	order_controller "go-mongodb-sample/app/controllers/order"
	product_controller "go-mongodb-sample/app/controllers/product"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBサーバーへの接続文字列
const connectionString = "mongodb://localhost:27017"

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("環境変数の読込に失敗しました: %v\r\n", err)
	}

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
	dbname := os.Getenv("DATABASE_NAME")
	customerController := customer_controller.NewCostumerController(connectionString, dbname, "customer")
	e.POST("/customer", customerController.Create)
	productController := product_controller.NewProductController(connectionString, dbname, "product")
	e.GET("/product/:id", productController.FindOne)
	e.POST("/product", productController.Create)
	orderController := order_controller.NewOrderController(connectionString, dbname, "order")
	e.GET("/orders/:customer_id", orderController.FindByCustomerID)
	e.POST("/order", orderController.Create)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}
