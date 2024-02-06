package product_controller

import (
	"context"
	"go-mongodb-sample/app/infrastructure"
	"go-mongodb-sample/app/infrastructure/product_infrastructure"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (pc ProductController) FindOne(c echo.Context) error {
	ID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// コンテキストを設定
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbm, err := infrastructure.NewMongoDBManager(ctx, pc.ConnectionString)
	if err != nil {
		return errors.Wrap(err, "NewMongoDBManager")
	}
	defer dbm.Client.Disconnect(ctx)

	pi := product_infrastructure.NewProductRepository(ctx, dbm.Client.Database(pc.DBName))
	order, err := pi.FindOne(ID)
	if errors.Is(err, product_infrastructure.ErrProductNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}
