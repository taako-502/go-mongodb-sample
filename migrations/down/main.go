package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/taako-502/go-mongodb-sample/migrations"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("環境変数の読込に失敗しました: %v\r\n", err)
	}

	uri := os.Getenv("MONGO_CONNECTION_STRING_FOR_MIGRATION")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	dbname := os.Getenv("DATABASE_NAME")
	db := client.Database(dbname)
	migrate.SetDatabase(db)

	one := 1 // 1つ前のバージョンに戻す
	if err := migrate.Down(context.Background(), one); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration completed")
}
