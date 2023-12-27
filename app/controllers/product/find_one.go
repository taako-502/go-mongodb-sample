package customer_controller

import (
	"context"
	"errors"
	product_infrastructure "go-mongodb-sample/app/infrastructures/products"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (pc ProductController) FindOne(c echo.Context) error {
	ID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(pc.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(pc.DBName).Collection(pc.CollectionName)
	pi := product_infrastructure.NewProduct(ctx, collection)
	order, err := pi.FindOne(ID)
	if errors.Is(err, product_infrastructure.ErrProductNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
